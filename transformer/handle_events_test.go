package transformer

import (
	"context"
	"testing"

	"github.com/jakoblorz/f1-metrics-transformer/packets"
	"github.com/stretchr/testify/assert"
)

func TestHandleEvents(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan int)
	s := HandleEvents(ctx, EventHandler{
		OnMotionPacket: func(pmd *packets.PacketMotionData) {
			assert.Equal(t, &motionData, pmd)
			close(sig)
		},
		OnSessionPacket: func(psd *packets.PacketSessionData) {
			t.Fail()
		},
	})

	readCh := make(chan interface{}, 1)
	writeCh := make(chan interface{}, 1)
	s.ReadFrom(readCh)
	s.WriteTo(writeCh)

	go s.Process()

	readCh <- &motionData
	assert.Equal(t, &motionData, <-writeCh)
	<-sig
}
