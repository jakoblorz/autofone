package cmd

import (
	"context"
	"fmt"
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/packets/process/reader"
	"github.com/jakoblorz/autofone/packets/process/writer"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/jakoblorz/autofone/pkg/privateapi"
	"github.com/jakoblorz/autofone/pkg/streamdb"
	"github.com/spf13/cobra"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type snapshotWriter struct {
	privateapi.Client
}

func (s *snapshotWriter) Write(reader io.Reader) (err error) {
	var r *privateapi.SnapshotCreateResponse
	r, err = s.Snapshots().CreateAndWrite(reader)
	if err == nil {
		file := strings.Split(r.File, "/")
		if len(file) != 3 {
			file = []string{"", "", ""}
		}
		log.Printf("successfully uploaded snapshot (%s/<redacted>/%s)", file[0], file[2])
	}
	return
}

var (
	url     string
	udp     int
	tcp     int
	filter  []uint
	logJSON bool
	logPack bool
	logRaw  bool

	token string

	devMode bool

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

			baseURL := "https://api.autofone.jakoblorz.de"
			if devMode {
				baseURL = "http://localhost:8080"
			}
			api := privateapi.New(token, baseURL)
			defer api.Close()

			db, err := streamdb.Open("autofone", &snapshotWriter{api}, streamdb.DebounceModeDelay)
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer db.Close()

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

			receivers := writer.NewReceiverRegistry()
			if tcp > 0 {
				http.Handle("/", websocket.Handler(receivers.Handle))
				go func() {
					log.Printf("awaiting connections from 0.0.0.0:%d", tcp)
					err := http.ListenAndServe(fmt.Sprintf(":%d", tcp), nil)
					if err != nil {
						log.Printf("%+v", err)
						sig <- os.Interrupt
						return
					}
				}()
			}

			log.Printf("awaiting packets from %s", conn.LocalAddr().String())
			stream := process.P{
				Context:  ctx,
				Hostname: host,
				C:        make(chan *process.M),
				S:        make(chan *process.M),
			}
			go func() {
				defer close(stream.C)
				udpr := reader.UDP{
					P:       &stream,
					LogPack: logPack,
					LogRaw:  logRaw,
					Verbose: verbose,
				}
				(&udpr).Read(ctx, conn, filter)
			}()
			go func() {
				writers := []writer.Writer{
					&writer.Websocket{
						P:                &stream,
						ReceiverRegistry: &receivers,
					},
					&writer.HTTP{
						P:       &stream,
						URL:     url,
						LogJSON: logJSON,
						Verbose: verbose,
					},
				}
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
						for _, w := range writers {
							go func(w writer.Writer) {
								defer func() {
									if r := recover(); r != nil {
										log.Printf("recovered from panic: %+v", r)
									}
								}()
								w.Write(m)
							}(w)
						}
					}
				}
			}()
			go func() {
				boltWriter := &writer.Bolt{
					P:  &stream,
					DB: db,
				}
				boltWriter.Motion = writer.NewMotionDebouncer(boltWriter, 0)
				boltWriter.Lap = writer.NewPacketDebouncer(boltWriter, 0)
				boltWriter.CarTelemetry = writer.NewPacketDebouncer(boltWriter, 0)
				boltWriter.CarStatus = writer.NewPacketDebouncer(boltWriter, 0)
				boltWriter.SessionHistory = writer.NewSessionHistoryDebouncer(boltWriter, 0)
				boltWriter.TyreSets = writer.NewTyreSetsDebouncer(boltWriter, 0)
				defer boltWriter.Close()

				for {
					select {
					case <-ctx.Done():
						return
					case m := <-stream.S:
						go func(w writer.Writer) {
							defer func() {
								if r := recover(); r != nil {
									log.Printf("recovered from panic: %+v", r)
								}
							}()
							w.Write(m)
						}(boltWriter)
					}
				}
			}()

			signal.Notify(sig, os.Interrupt)
			<-sig
		},
	}
)

func init() {
	rootCmd.AddCommand(bindCmd)

	// settings
	bindCmd.Flags().UintSliceVar(&filter, "filter", []uint{uint(constants.PacketFinalClassification) /*, uint(constants.PacketMotion) */}, "Filter the packets that are to be relayed, no filter means accepting all")

	// io
	bindCmd.Flags().IntVar(&udp, "udp", 20777, "UDP port to listen on; 20777 is the F1 2021/2022 default UDP port")
	bindCmd.Flags().IntVar(&tcp, "tcp", -1, "TCP port to listen on for websocket connections; -1 means websocket is disabled")
	bindCmd.Flags().StringVar(&url, "http", "https://localhost:8081/f1", "FQURL to post the packets to; if empty, no request is sent")

	// auth
	bindCmd.Flags().StringVar(&token, "token", "", "Token to authenticate against autofone.jakoblorz.de")

	// logging flags
	bindCmd.Flags().BoolVar(&logJSON, "json", false, "Log JSON sent to the HTTP Server")
	bindCmd.Flags().BoolVar(&logPack, "pack", false, "Log unpacked packets")
	bindCmd.Flags().BoolVar(&logRaw, "bytes", false, "Log bytes received from the UDP socket")

	bindCmd.Flags().BoolVar(&devMode, "dev", false, "Enable development mode")
}
