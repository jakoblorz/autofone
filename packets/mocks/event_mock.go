package mocks

import (
	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/constants/event"
	"github.com/jakoblorz/autofone/packets"
)

var (
	PacketEventButtons21Bytes = []byte{229, 7, 1, 5, 1, 3, 187, 134, 38, 178, 108, 178, 251, 17, 189, 180, 189, 68, 166, 123, 0, 0, 19, 255, 66, 85, 84, 78, 4, 0, 0, 0, 24, 0, 0, 0}
	PacketEventButtons21      = packets.PacketEventButtons21{
		Header: packets.PacketHeader21{
			PacketFormat:            constants.PacketFormat_2021,
			GameMajorVersion:        1,
			GameMinorVersion:        5,
			PacketVersion:           1,
			PacketID:                3,
			SessionTime:             1517.6480712890625,
			SessionUID:              1295825497714230971,
			FrameIdentifier:         31654,
			PlayerCarIndex:          19,
			SecondaryPlayerCarIndex: 255,
		},
		EventStringCode: [4]uint8{
			[]byte(event.ButtonStatus)[0],
			[]byte(event.ButtonStatus)[1],
			[]byte(event.ButtonStatus)[2],
			[]byte(event.ButtonStatus)[3],
		},
		EventDetails: packets.Buttons21{
			ButtonStatus: 4,
		},
	}
)
