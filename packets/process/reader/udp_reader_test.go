package reader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/udptest"
	"github.com/stretchr/testify/assert"
)

func TestUDP_read(t *testing.T) {
	type testCase struct {
		name string
		data []byte
		exp  interface{}
	}
	createTestCase := func(name string, data []byte, exp interface{}) testCase {
		return testCase{
			name: fmt.Sprintf("should decode and unpack %s from udp connection", name),
			data: data,
			exp:  exp,
		}
	}
	tests := []testCase{
		createTestCase("PacketCarSetupData21", mocks.PacketCarSetupData21Bytes, &mocks.PacketCarSetupData21),
		createTestCase("PacketCarStatusData21", mocks.PacketCarStatusData21Bytes, &mocks.PacketCarStatusData21),
		createTestCase("PacketCarTelemetryData21", mocks.PacketCarTelemetryData21Bytes, &mocks.PacketCarTelemetryData21),
		createTestCase("PacketEventButtons21", mocks.PacketEventButtons21Bytes, &mocks.PacketEventButtons21),
		createTestCase("PacketLapData21", mocks.PacketLapData21Bytes, &mocks.PacketLapData21),
		createTestCase("PacketMotionData21", mocks.PacketMotionData21Bytes, &mocks.PacketMotionData21),
		createTestCase("PacketParticipantsData21", mocks.PacketParticipantsData21Bytes, &mocks.PacketParticipantsData21),
		createTestCase("PacketSessionData21", mocks.PacketSessionData21Bytes, &mocks.PacketSessionData21),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			clientCh := make(chan net.Conn)
			go func() {
				defer close(clientCh)
				io.Copy(<-clientCh, bytes.NewReader(tt.data))
			}()

			stream := process.P{
				Context: ctx,
				C:       make(chan *process.M, 1),
			}
			r := UDP{
				P: &stream,
			}
			if !assert.NoError(t, udptest.NewConn(func(clientConn, serverConn net.Conn) error {
				clientCh <- clientConn
				(&r).Read(ctx, serverConn.(*net.UDPConn), []uint{})
				return nil
			}, 2*time.Second), "should be able to read from udp connection") {
				t.Fail()
				return
			}

			if !assert.Equal(t, tt.exp, (<-stream.C).Pack, "should decode and unpack packet") {
				t.Fail()
				return
			}
		})
	}
}
