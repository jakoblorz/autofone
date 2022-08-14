package reader

import (
	"context"
	"fmt"
	"net"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/log"
)

type UDP struct {
	*process.P

	Verbose bool
	LogRaw  bool
	LogPack bool
}

func (ch *UDP) Read(ctx context.Context, conn *net.UDPConn, filter []uint) {

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
		if err = packets.Read_LE(buf, header); err != nil {
			log.Printf("%+v", err)
			continue
		}

		if ch.Verbose || ch.LogRaw {
			message := fmt.Sprintf("received %d bytes, representing packet %d -> proceed", n, header.PacketID)
			if ch.LogRaw {
				message = fmt.Sprintf("%s: %+b", message, buf)
			}
			log.Print(message)
		}

		pack := packets.ByPacketID(header.PacketID, header.PacketFormat)
		if pack == nil {
			log.Printf("invalid packet: %d", header.PacketID)
			continue
		}

		if err = packets.Read_LE(buf, pack); err != nil {
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
			if err = packets.Read_LE(buf, pack); err != nil {
				log.Printf("failed to read event packet %d: %+v", header.PacketID, err)
				continue
			}
		}

		var c uint8
		for _, f := range filter {
			c = uint8(f) ^ header.PacketID
			if c == 0 {
				break
			}
		}
		if c == 0 {
			if ch.LogPack {
				log.Printf("processing package: %+v", pack)
			}

			ch.C <- &process.M{
				Header: *header,
				Pack:   pack,
				Buffer: buf,
			}
		}
		if ch.S != nil {
			ch.S <- &process.M{
				Header: *header,
				Pack:   pack,
				Buffer: buf,
			}
		}
	}
}
