package pipe

import (
	"context"

	"github.com/jakoblorz/metrikxd/packets"
	"github.com/jakoblorz/metrikxd/pkg/step"
)

type EventHandler struct {
	OnMotionPacket              func(*packets.PacketMotionData)
	OnSessionPacket             func(*packets.PacketSessionData)
	OnLapPacket                 func(*packets.PacketLapData)
	OnEventPacket               func(*packets.PacketEventData)
	OnParticipantsPacket        func(*packets.PacketParticipantsData)
	OnCarSetupPacket            func(*packets.PacketCarSetupData)
	OnCarTelemetryPacket        func(*packets.PacketCarTelemetryData)
	OnCarStatusPacket           func(*packets.PacketCarStatusData)
	OnFinalClassificationPacket func(*packets.PacketFinalClassificationData)
	OnLobbyInformationPacket    func(*packets.PacketLobbyInfoData)
	OnCarDamagePacket           func(*packets.PacketCarDamageData)
	OnSessionHistoryPacket      func(*packets.PacketSessionHistoryData)
}

func HandleEvents(ctx context.Context, handler EventHandler) step.Step {
	return Splicer(ctx, func(pack interface{}) {
		switch p := pack.(type) {
		case *packets.PacketMotionData:
			if handler.OnMotionPacket != nil {
				handler.OnMotionPacket(p)
			}
		case *packets.PacketSessionData:
			if handler.OnSessionPacket != nil {
				handler.OnSessionPacket(p)
			}
		case *packets.PacketLapData:
			if handler.OnLapPacket != nil {
				handler.OnLapPacket(p)
			}
		case *packets.PacketEventData:
			if handler.OnEventPacket != nil {
				handler.OnEventPacket(p)
			}
		case *packets.PacketParticipantsData:
			if handler.OnParticipantsPacket != nil {
				handler.OnParticipantsPacket(p)
			}
		case *packets.PacketCarSetupData:
			if handler.OnCarSetupPacket != nil {
				handler.OnCarSetupPacket(p)
			}
		case *packets.PacketCarTelemetryData:
			if handler.OnCarTelemetryPacket != nil {
				handler.OnCarTelemetryPacket(p)
			}
		case *packets.PacketCarStatusData:
			if handler.OnCarStatusPacket != nil {
				handler.OnCarStatusPacket(p)
			}
		case *packets.PacketFinalClassificationData:
			if handler.OnFinalClassificationPacket != nil {
				handler.OnFinalClassificationPacket(p)
			}
		case *packets.PacketLobbyInfoData:
			if handler.OnLobbyInformationPacket != nil {
				handler.OnLobbyInformationPacket(p)
			}
		case *packets.PacketCarDamageData:
			if handler.OnCarDamagePacket != nil {
				handler.OnCarDamagePacket(p)
			}
		case *packets.PacketSessionHistoryData:
			if handler.OnSessionHistoryPacket != nil {
				handler.OnSessionHistoryPacket(p)
			}
		}
	})
}
