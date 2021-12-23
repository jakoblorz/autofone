package pipe

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"unsafe"

	"github.com/jakoblorz/f1-metrics-transformer/constants"
	"github.com/jakoblorz/f1-metrics-transformer/constants/event"
	"github.com/jakoblorz/f1-metrics-transformer/packets"
	"github.com/jakoblorz/f1-metrics-transformer/pkg/log"
	"github.com/jakoblorz/f1-metrics-transformer/pkg/step"
)

type PacketReader struct {
	step.Step
	io.Reader

	filter    []uint
	logBytes  bool
	logStruct bool
}

func (u *PacketReader) read(ch chan<- interface{}) {
	for {
		buf := make([]byte, 1024+1024/2)
		n, err := u.Read(buf)
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
		for _, f := range u.filter {
			c = uint8(f) ^ header.PacketID
			if c == 0 {
				break
			}
		}
		if c != 0 {
			log.Verbosef("received %d bytes, representing packet %d -> dropping", n, header.PacketID)
			continue
		}

		message := fmt.Sprintf("received %d bytes, representing packet %d -> proceed", n, header.PacketID)
		if u.logBytes {
			message = fmt.Sprintf("%s: %+b", message, buf)
		}
		log.Print(message)

		pack := newPacketById(header.PacketID)
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

		if u.logStruct {
			log.Printf("decoded package %d: %+v", header.PacketID, pack)
		}

		ch <- pack
	}
}

type PacketReaderOptions struct {
	Filter           []uint
	LogIncomingBytes bool
	LogDecodedStruct bool
}

func ReadUDPPackets(ctx context.Context, conn net.Conn, opts *PacketReaderOptions) *PacketReader {
	r := &PacketReader{Reader: conn, filter: opts.Filter, logBytes: opts.LogIncomingBytes, logStruct: opts.LogDecodedStruct}
	r.Step = step.Emitter(ctx, r.read)
	return r
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
	case constants.PacketMotion:
		return new(packets.PacketMotionData)
	case constants.PacketSession:
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
		return new(packets.PacketFinalClassificationData)
	case constants.PacketLobbyInfo:
		return new(packets.PacketLobbyInfoData)
	case constants.PacketCarDamage:
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
		return new(packets.SpeedTrap)
	}

	return nil
}
