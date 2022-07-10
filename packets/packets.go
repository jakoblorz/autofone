package packets

import (
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/constants/event"
)

func ByPacketID(packetId uint8, packetFormat uint16) interface{} {
	switch packetId {
	case constants.PacketMotion:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketMotionData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketMotionData21)
		}
	case constants.PacketSession:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketSessionData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketSessionData21)
		}
	case constants.PacketLap:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketLapData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketLapData21)
		}
	case constants.PacketParticipants:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketCarDamageData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketCarDamageData21)
		}
	case constants.PacketCarSetup:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketCarSetupData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketCarSetupData21)
		}
	case constants.PacketCarTelemetry:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketCarTelemetryData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketCarTelemetryData21)
		}
	case constants.PacketCarStatus:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketCarStatusData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketCarStatusData21)
		}
	case constants.PacketFinalClassification:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketFinalClassificationData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketFinalClassificationData21)
		}
	case constants.PacketLobbyInfo:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketLobbyInfoData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketLobbyInfoData21)
		}
	case constants.PacketCarDamage:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketCarDamageData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketCarDamageData21)
		}
	case constants.PacketSessionHistory:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketSessionHistoryData22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketSessionHistoryData21)
		}
	case constants.PacketEvent:
		return new(PacketEventHeader)
	}

	return nil
}

func ByEventHeader(h *PacketEventHeader, packetFormat uint16) interface{} {
	switch h.EventCodeString() {
	case event.FastestLap:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventFastestLap22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventFastestLap21)
		}
	case event.SpeedTrapTriggered:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventSpeedTrap22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventSpeedTrap21)
		}
	case event.PenaltyIssued:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventPenalty22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventPenalty21)
		}
	case event.Flashback:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventFlashback22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventFlashback21)
		}
	case event.StartLights:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventStartLights22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventStartLights21)
		}
	case event.ButtonStatus:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventButtons22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventButtons21)
		}

	case event.Retirement:
	case event.TeamMateInPit:
	case event.RaceWinner:
	case event.DriveThroughServed:
	case event.StopGoServed:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventGenericVehicleEvent22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventGenericVehicleEvent21)
		}

	case event.SessionStarted:
	case event.SessionEnded:
	case event.DRSEnabled:
	case event.DRSDisabled:
	case event.ChequeredFlag:
	case event.LightsOut:
		if packetFormat == constants.PacketFormat_2022 {
			return new(PacketEventGenericSessionEvent22)
		}
		if packetFormat == constants.PacketFormat_2021 {
			return new(PacketEventGenericSessionEvent21)
		}
	}

	return nil
}
