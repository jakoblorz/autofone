package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"unsafe"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/spf13/cobra"
)

var (
	to     string
	port   int
	filter []uint

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

See https://github.com/anilmisirlioglu/f1-telemetry-go/blob/master/pkg/env/packet_ids.go
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

			ch := make(chan interface{})
			go func(conn *net.UDPConn) {
				defer close(ch)

			READ_UDP:
				for {
					buf := make([]byte, 1024+1024/2)
					_, _, err := conn.ReadFromUDP(buf)
					if err != nil {
						log.Printf("read error: %+v", err)
						continue
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
						continue READ_UDP
					}

					pack := newPacketById(header.PacketID)
					if pack == nil {
						log.Printf("invalid packet: %d", header.PacketID)
						continue
					}

					if err = read(buf, pack); err != nil {
						log.Printf("failed to read packet %d: %+v", header.PacketID, err)
						continue
					}

					if header.PacketID == env.PacketEvent {
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
					ch <- pack
				}
			}(conn)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go func() {
				onProxyPacket := func(pack interface{}) {
					defer func() {
						if err := recover(); err != nil {
							log.Printf("%+v", err)
							err = binary.Write(os.Stderr, binary.LittleEndian, pack)
							if err != nil {
								log.Printf("%+v", err)
							}
						}
					}()
					data, err := json.Marshal(pack)
					if err != nil {
						panic(err)
					}
					req, err := http.NewRequest("POST", to, bytes.NewBuffer(data))
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
						onProxyPacket(pack)
					}
				}
			}()

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig
		},
	}
)

func init() {
	rootCmd.AddCommand(bindCmd)

	bindCmd.Flags().StringVar(&to, "to", "https://localhost:8081/f1", "FQURL to post the packets to")
	bindCmd.Flags().IntVar(&port, "port", 20777, "UDP port to listen on")
	bindCmd.Flags().UintSliceVar(&filter, "filter", []uint{uint(env.PacketFinalClassification)}, "Filter the packets that are to be relayed, no filter means accepting all")
}

func read(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func newPacketById(packetId uint8) interface{} {
	switch packetId {
	case env.PacketMotion:
		return new(packets.PacketMotionData)
	case env.PacketSession:
		return new(packets.PacketSessionData)
	case env.PacketLap:
		return new(packets.PacketLapData)
	case env.PacketEvent:
		return new(packets.PrePacketEventData)
	case env.PacketParticipants:
		return new(packets.PacketParticipantsData)
	case env.PacketCarSetup:
		return new(packets.PacketCarSetupData)
	case env.PacketCarTelemetry:
		return new(packets.PacketCarTelemetryData)
	case env.PacketCarStatus:
		return new(packets.PacketCarStatusData)
	case env.PacketFinalClassification:
		return new(packets.PacketFinalClassificationData)
	case env.PacketLobbyInfo:
		return new(packets.PacketLobbyInfoData)
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
		return new(packets.SpeedTrap)
	}

	return nil
}
