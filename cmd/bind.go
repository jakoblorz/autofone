package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/packets/sql"
	"github.com/jakoblorz/autofone/pkg/gcs"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/jakoblorz/autofone/pkg/streamdb"
	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
)

var (
	url     string
	udp     int
	tcp     int
	filter  []uint
	logJSON bool
	logPack bool
	logRaw  bool

	devMode bool

	socketPool = &socketHandler{
		RWMutex: new(sync.RWMutex),
		Source:  rand.NewSource(time.Now().UnixNano()),
		conns:   make(map[string]*websocket.Conn),
	}

	clientPool = &sync.Pool{
		New: func() interface{} {
			return &http.Client{}
		},
	}

	bindCmd = &cobra.Command{
		Use:   "bind",
		Short: "bind to the F1 UDP port and stream all selected packets to a destination HTTP server",
		Long: `	
binds to the localhost F1 UDP port. Received packets are processed 
and then sent to the provided HTTP server via POST. 
Use --filter to select the packets to subscribe to, otherwise 
make sure that the HTTP server's DDOS protection is not kicking in.

See https://github.com/anilmisirlioglu/f1-telemetry-go/blob/master/pkg/constants/packet_ids.go
for the packet ids to select.
`,
		Run: func(cmd *cobra.Command, args []string) {
			sig := make(chan os.Signal, 1)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			conn, err := net.ListenUDP("udp", &net.UDPAddr{
				IP:   net.ParseIP("localhost"),
				Port: udp,
			})
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer conn.Close()

			if tcp > 0 {
				http.Handle("/", websocket.Handler(socketPool.handleConn))
				go func() {
					log.Verbosef("awaiting connections from 0.0.0.0:%d", tcp)
					err := http.ListenAndServe(fmt.Sprintf(":%d", tcp), nil)
					if err != nil {
						log.Printf("%+v", err)
						sig <- os.Interrupt
						return
					}
				}()
			}

			replica := gcs.NewReplicaClient()
			replica.Path = mac
			replica.Bucket = storageBucket

			db, err := new(streamdb.I).Replicated(ctx, "autofone.sqlite3", replica.WithClient(storageClient))
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer db.Close()

			err = sql.Init(db.DB)
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			db.MustHardSync(ctx)

			log.Verbosef("awaiting packets from %s", conn.LocalAddr().String())
			stream := process.P{
				Context:   ctx,
				Hostname:  host,
				SessionID: sessionID,
				C:         make(chan *process.M),
			}
			go func() {
				defer close(stream.C)

				r := reader(stream)
				(&r).read(ctx, conn, filter)
			}()
			go func() {
				var (
					dbwriter  = sqlwriter(stream)
					urlwriter = httpwriter(stream)
				)
				for {
					select {
					case <-ctx.Done():
						for {
							_, ok := <-stream.C
							if !ok {
								return
							}
						}
					case m := <-stream.C:
						go (&dbwriter).write(m, db)
						go (&urlwriter).write(m, url)
						go socketPool.write(m.Pack)
					}
				}
			}()

			signal.Notify(sig, os.Interrupt)
			<-sig
		},
	}
)

type reader process.P

func (ch *reader) read(ctx context.Context, conn *net.UDPConn, filter []uint) {

READ_UDP:
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		buf := make([]byte, 1024+1024/2)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("read error: %+v", err)
			return
		}

		header := new(packets.PacketHeader)
		if err = read(buf, header); err != nil {
			log.Printf("%+v", err)
			continue
		}

		var c uint8
		for _, f := range filter {
			c = uint8(f) ^ header.PacketID
			if c == 0 {
				break
			}
		}
		if c != 0 {
			log.Verbosef("received %d bytes, representing packet %d -> dropping", n, header.PacketID)
			continue READ_UDP
		} else if verbose || logRaw {
			message := fmt.Sprintf("received %d bytes, representing packet %d -> proceed", n, header.PacketID)
			if logRaw {
				message = fmt.Sprintf("%s: %+b", message, buf)
			}
			log.Print(message)
		}

		pack := packets.ByPacketID(header.PacketID, header.PacketFormat)
		if pack == nil {
			log.Printf("invalid packet: %d", header.PacketID)
			continue
		}

		if err = read(buf, pack); err != nil {
			log.Printf("failed to read packet %d: %+v", header.PacketID, err)
			continue
		}

		if header.PacketID == constants.PacketEvent {
			h := pack.(*packets.PacketEventHeader)
			pack = packets.ByEventHeader(h, header.PacketFormat)
			if pack == nil {
				log.Printf("invalid event packet: %d", header.PacketID)
				continue
			}
			if err = read(buf, pack); err != nil {
				log.Printf("failed to read event packet %d: %+v", header.PacketID, err)
				continue
			}
		}

		if logPack {
			log.Printf("processing package: %+v", pack)
		}

		ch.C <- &process.M{
			Header: *header,
			Pack:   pack,
			Buffer: buf,
		}

	}
}

type sqlwriter process.P

func (ch *sqlwriter) write(m *process.M, db *streamdb.I) {
	tx, err := db.Beginx()
	if err != nil {
		log.Printf("tx begin() error: %+v", err)
		return
	}
	defer tx.Rollback()

	err = (&sql.Packet{
		Hostname:     ch.Hostname,
		PacketHeader: m.Header,
		Data:         m.Buffer,
	}).Write(ch.Context, tx)
	if err != nil {
		log.Printf("tx write() error: %+v", err)
		return
	}

	err = db.SoftSync(ch.Context)
	if err != nil {
		log.Printf("tx sync(1) error: %+v", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("tx commit() error: %+v", err)
		return
	}

	err = db.SoftSync(ch.Context)
	if err != nil {
		log.Printf("tx sync(2) error: %+v", err)
		return
	}
}

type httpwriter process.P

func (*httpwriter) write(m *process.M, to string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
			err = binary.Write(os.Stderr, binary.LittleEndian, m.Pack)
			if err != nil {
				log.Printf("%+v", err)
			}
		}
	}()
	data, err := json.Marshal(m.Pack)
	if err != nil {
		panic(err)
	}

	if verbose || logJSON {
		message := fmt.Sprintf("posting with len = %d bytes json payload", len(data))
		if logJSON {
			message = fmt.Sprintf("%s: %s", message, string(data))
		}
		log.Print(message)
	}

	if len(to) == 0 {
		return
	}

	req, err := http.NewRequest("POST", strings.ReplaceAll(to, "{{packetID}}", fmt.Sprintf("%d", m.Header.PacketID)), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := clientPool.Get().(*http.Client)
	res, err := client.Do(req)
	clientPool.Put(client)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		panic(err)
	}
}

type socketHandler struct {
	*sync.RWMutex
	rand.Source
	conns map[string]*websocket.Conn
}

func (s *socketHandler) write(raw interface{}) {
	failedHandles := make([]string, 0)

	s.RLock()
	for handle, conn := range s.conns {
		err := websocket.JSON.Send(conn, raw)
		if err != nil {
			log.Printf("%+v", err)
			failedHandles = append(failedHandles, handle)
		}
	}
	s.RUnlock()

	for _, handle := range failedHandles {
		s.unregisterConn(handle)
	}
}

func (s *socketHandler) registerConn(conn *websocket.Conn) string {
	handle := fmt.Sprintf("%s-%f", conn.RemoteAddr(), float64(time.Now().UnixNano())*rand.New(s.Source).Float64())
	s.Lock()
	s.conns[handle] = conn
	s.Unlock()
	return handle
}

func (s *socketHandler) unregisterConn(handle string) error {
	s.Lock()
	conn, ok := s.conns[handle]
	if !ok {
		s.Unlock()
		return nil
	}
	delete(s.conns, handle)
	s.Unlock()
	return conn.Close()
}

func (s *socketHandler) handleConn(ws *websocket.Conn) {
	handle := s.registerConn(ws)
	defer s.unregisterConn(handle)
	for {
		// discard all messages
		_, err := io.Copy(ioutil.Discard, ws)
		if err != nil {
			log.Printf("%+v", err)
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(bindCmd)

	// settings
	bindCmd.Flags().UintSliceVar(&filter, "filter", []uint{uint(constants.PacketFinalClassification)}, "Filter the packets that are to be relayed, no filter means accepting all")

	// io
	bindCmd.Flags().IntVar(&udp, "udp", 20777, "UDP port to listen on; 20777 is the F1 2021/2022 default UDP port")
	bindCmd.Flags().IntVar(&tcp, "tcp", -1, "TCP port to listen on for websocket connections; -1 means websocket is disabled")
	bindCmd.Flags().StringVar(&url, "http", "https://localhost:8081/f1", "FQURL to post the packets to; if empty, no request is sent")

	// logging flags
	bindCmd.Flags().BoolVar(&logJSON, "json", false, "Log JSON sent to the HTTP Server")
	bindCmd.Flags().BoolVar(&logPack, "pack", false, "Log unpacked packets")
	bindCmd.Flags().BoolVar(&logRaw, "bytes", false, "Log bytes received from the UDP socket")

	bindCmd.Flags().BoolVar(&devMode, "dev", false, "Enable development mode")
}

func read(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}
