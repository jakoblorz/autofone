package state_set

import (
	"github.com/jakoblorz/metrikxd/packets"
	ps "github.com/jakoblorz/metrikxd/pkg/pub_sub"
)

type SessionState interface {
	ps.SubscriptionReceiver
	OnLobbyInfoDataReceived(*packets.PacketLobbyInfoData)
}
