package session

import (
	"sync"

	"github.com/jakoblorz/metrikxd/packets"
	ps "github.com/jakoblorz/metrikxd/pkg/pub_sub"
	ss "github.com/jakoblorz/metrikxd/state_set"
)

var Instance ss.SessionState = &sessionStateImpl{
	Locker:                       *new(sync.Locker),
	SubscriptionReceiverNotifier: ps.NewReceiverNotifier(),
}

type sessionStateImpl struct {
	sync.Locker
	ps.SubscriptionReceiverNotifier

	lobbyInfo *packets.PacketLobbyInfoData
}

func (s *sessionStateImpl) OnLobbyInfoDataReceived(li *packets.PacketLobbyInfoData) {
	s.Lock()
	defer s.Unlock()

	s.lobbyInfo = li
	s.Notify()
}
