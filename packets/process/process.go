package process

import (
	"context"

	"github.com/jakoblorz/autofone/packets"
)

type M struct {
	Header packets.PacketHeader
	Pack   interface{}
	Buffer []byte
}

type P struct {
	context.Context

	Hostname string
	DeviceID string
	UserID   string

	C chan *M
	S chan *M
}
