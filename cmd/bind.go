package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"unsafe"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/constants/event"
	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/spf13/cobra"
)

var (
	post    string
	port    int
	filter  []uint
	logJSON bool
	logPack bool
	logRaw  bool

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
			conn, err := net.ListenUDP("udp", &net.UDPAddr{
				IP:   net.ParseIP("localhost"),
				Port: port,
			})
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer conn.Close()

			log.Verbosef("awaiting packets from %s", conn.LocalAddr().String())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			stream := process(make(chan struct {
				header packets.PacketHeader
				raw    interface{}
			}))
			go func() {
				defer close(stream)
				stream.readPacketsFromConn(conn, filter)
			}()

			go stream.handlePackets(ctx, post)

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig
		},
	}
)

type process chan struct {
	header packets.PacketHeader
	raw    interface{}
}

func (ch process) readPacketsFromConn(conn *net.UDPConn, filter []uint) {

READ_UDP:
	for {
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

		pack := newPacketById(header.PacketID, fmt.Sprint(header.PacketFormat))
		if pack == nil {
			log.Printf("invalid packet: %d", header.PacketID)
			continue
		}

		if err = read(buf, pack); err != nil {
			log.Printf("failed to read packet %d: %+v", header.PacketID, err)
			continue
		}

		if header.PacketID == constants.PacketEvent {
			details := resolveEventDetails(pack.(*packets.PrePacketEventData))
			pre := pack.(*packets.PrePacketEventData)
			if details != nil {
				err = read(pre.EventDetails[:unsafe.Sizeof(details)], details)
				if err != nil {
					log.Printf("event packet details read error: %+v", err)
					continue
				}
			}
			pack = &packets.PacketEventData{
				Header:          pre.Header,
				EventStringCode: pre.EventStringCode,
				EventDetails:    details,
			}
		}

		if logPack {
			log.Printf("processing package: %+v", pack)
		}

		ch <- struct {
			header packets.PacketHeader
			raw    interface{}
		}{
			header: *header,
			raw:    pack,
		}

	}
}

func (ch process) handlePackets(ctx context.Context, to string) {
	for {
		select {
		case <-ctx.Done():
			for {
				_, ok := <-ch
				if !ok {
					return
				}
			}
		case pack := <-ch:
			ch.writePacketToHTTP(pack, to)
		}
	}
}

func (process) writePacketToHTTP(pack struct {
	header packets.PacketHeader
	raw    interface{}
}, to string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
			err = binary.Write(os.Stderr, binary.LittleEndian, pack)
			if err != nil {
				log.Printf("%+v", err)
			}
		}
	}()
	data, err := json.Marshal(pack.raw)
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

	req, err := http.NewRequest("POST", strings.ReplaceAll(to, "{{packetID}}", fmt.Sprintf("%d", pack.header.PacketID)), bytes.NewBuffer(data))
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

func init() {
	rootCmd.AddCommand(bindCmd)

	bindCmd.Flags().StringVar(&post, "to", "https://localhost:8081/f1", "FQURL to post the packets to; if empty, no request is sent")
	bindCmd.Flags().IntVar(&port, "port", 20777, "UDP port to listen on")
	bindCmd.Flags().UintSliceVar(&filter, "filter", []uint{uint(constants.PacketFinalClassification)}, "Filter the packets that are to be relayed, no filter means accepting all")
	bindCmd.Flags().BoolVar(&logJSON, "json", false, "Log JSON sent to destination")
	bindCmd.Flags().BoolVar(&logPack, "pack", false, "Log unmarshaled data in go representation")
	bindCmd.Flags().BoolVar(&logRaw, "bytes", false, "Log bytes received")
}

func read(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func newPacketById(packetId uint8, packetFormat string) interface{} {
	switch packetId {
	case constants.PacketMotion:
		return new(packets.PacketMotionData)
	case constants.PacketSession:
		if packetFormat == "2022" {
			return new(packets.PacketSessionData22)
		}
		return new(packets.PacketSessionData)
	case constants.PacketLap:
		return new(packets.PacketLapData)
	case constants.PacketEvent:
		return new(packets.PrePacketEventData)
	case constants.PacketParticipants:
		return new(packets.PacketParticipantsData)
	case constants.PacketCarSetup:
		return new(packets.PacketCarSetupData)
	case constants.PacketCarTelemetry:
		return new(packets.PacketCarTelemetryData)
	case constants.PacketCarStatus:
		return new(packets.PacketCarStatusData)
	case constants.PacketFinalClassification:
		if packetFormat == "2022" {
			return new(packets.PacketFinalClassificationData22)
		}
		return new(packets.PacketFinalClassificationData)
	case constants.PacketLobbyInfo:
		return new(packets.PacketLobbyInfoData)
	case constants.PacketCarDamage:
		if packetFormat == "2022" {
			return new(packets.PacketCarDamageData22)
		}
		return new(packets.PacketCarDamageData)
	case constants.PacketSessionHistory:
		return new(packets.PacketSessionHistoryData)
	}

	return nil
}

func resolveEventDetails(pre *packets.PrePacketEventData) interface{} {
	switch string(pre.EventStringCode[:]) {
	case event.FastestLap:
		return new(packets.FastestLap)
	case event.Retirement:
		return new(packets.Retirement)
	case event.TeamMateInPit:
		return new(packets.TeamMateInPits)
	case event.RaceWinner:
		return new(packets.RaceWinner)
	case event.PenaltyIssued:
		return new(packets.Penalty)
	case event.SpeedTrapTriggered:
		if fmt.Sprint(pre.Header.PacketFormat) == "2022" {
			return new(packets.SpeedTrapF22)
		}
		return new(packets.SpeedTrap)
	}

	return nil
}
