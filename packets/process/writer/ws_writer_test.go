package writer

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jakoblorz/autofone/packets"
	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/wstest"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
)

func TestWebsocket_Write(t *testing.T) {

	type testCase struct {
		name        string
		data        *process.M
		decoderFunc func() interface{}
	}
	createTestCase := func(name string, data interface{}, decoderFunc func() interface{}) testCase {
		return testCase{
			name:        fmt.Sprintf("should write packet %s to websocket, encoded as JSON", name),
			data:        &process.M{Pack: data},
			decoderFunc: decoderFunc,
		}
	}
	tests := []testCase{
		createTestCase("PacketCarSetupData21", &mocks.PacketCarSetupData21, func() interface{} {
			return new(packets.PacketCarSetupData21)
		}),
		createTestCase("PacketCarStatusData21", &mocks.PacketCarStatusData21, func() interface{} {
			return new(packets.PacketCarStatusData21)
		}),
		createTestCase("PacketCarTelemetryData21", &mocks.PacketCarTelemetryData21, func() interface{} {
			return new(packets.PacketCarTelemetryData21)
		}),
		createTestCase("PacketEventButtons21", &mocks.PacketEventButtons21, func() interface{} {
			return new(packets.PacketEventButtons21)
		}),
		createTestCase("PacketLapData21", &mocks.PacketLapData21, func() interface{} {
			return new(packets.PacketLapData21)
		}),
		createTestCase("PacketMotionData21", &mocks.PacketMotionData21, func() interface{} {
			return new(packets.PacketMotionData21)
		}),
		createTestCase("PacketParticipantsData21", &mocks.PacketParticipantsData21, func() interface{} {
			return new(packets.PacketParticipantsData21)
		}),
		createTestCase("PacketSessionData21", &mocks.PacketSessionData21, func() interface{} {
			return new(packets.PacketSessionData21)
		}),
	}

	// dirty tricks all in this one, don't look
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := make(chan int)
			registry := NewReceiverRegistry()

			wg := new(sync.WaitGroup)
			wg.Add(1)

			go func(registry *ReceiverRegistry) {
				wstest.NewConn(func(clientConn, serverConn *websocket.Conn) error {
					defer close(sig)

					go registry.Handle(serverConn)
					go func() {
						<-time.After(1 * time.Second)
						wg.Done()
					}()

					decoded := tt.decoderFunc()
					err := websocket.JSON.Receive(clientConn, decoded)
					if !assert.NoError(t, err) {
						t.Fail()
						return nil
					}

					if !assert.Equal(t, tt.data.Pack, decoded) {
						t.Fail()
						return nil
					}
					return nil
				})
			}(&registry)

			wg.Wait()
			(&Websocket{
				ReceiverRegistry: &registry,
			}).Write(tt.data)
			<-sig
		})
	}
}
