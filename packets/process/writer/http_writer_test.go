package writer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jakoblorz/autofone/packets"

	"github.com/jakoblorz/autofone/constants"
	"github.com/jakoblorz/autofone/packets/mocks"
	"github.com/jakoblorz/autofone/packets/process"
)

func TestHTTP_Write(t *testing.T) {

	type testCase struct {
		name        string
		data        interface{}
		decoderFunc func() interface{}
		packetID    uint8
		path        string
	}
	createTestCase := func(name string, packetID uint8, data interface{}, decoderFunc func() interface{}) testCase {
		return testCase{
			name:        fmt.Sprintf("should POST packet %s to /%d, encoded as JSON", name, packetID),
			data:        data,
			path:        fmt.Sprintf("/%d", packetID),
			packetID:    packetID,
			decoderFunc: decoderFunc,
		}
	}
	tests := []testCase{
		createTestCase("PacketCarSetupData21", constants.PacketCarSetup, &mocks.PacketCarSetupData21, func() interface{} {
			return new(packets.PacketCarSetupData21)
		}),
		createTestCase("PacketCarStatusData21", constants.PacketCarStatus, &mocks.PacketCarStatusData21, func() interface{} {
			return new(packets.PacketCarStatusData21)
		}),
		createTestCase("PacketCarTelemetryData21", constants.PacketCarTelemetry, &mocks.PacketCarTelemetryData21, func() interface{} {
			return new(packets.PacketCarTelemetryData21)
		}),
		createTestCase("PacketEventButtons21", constants.PacketEvent, &mocks.PacketEventButtons21, func() interface{} {
			return new(packets.PacketEventButtons21)
		}),
		createTestCase("PacketLapData21", constants.PacketLap, &mocks.PacketLapData21, func() interface{} {
			return new(packets.PacketLapData21)
		}),
		createTestCase("PacketMotionData21", constants.PacketMotion, &mocks.PacketMotionData21, func() interface{} {
			return new(packets.PacketMotionData21)
		}),
		createTestCase("PacketParticipantsData21", constants.PacketParticipants, &mocks.PacketParticipantsData21, func() interface{} {
			return new(packets.PacketParticipantsData21)
		}),
		createTestCase("PacketSessionData21", constants.PacketSession, &mocks.PacketSessionData21, func() interface{} {
			return new(packets.PacketSessionData21)
		}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := make(chan int)
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer close(sig)

				decoded := tt.decoderFunc()
				err := json.NewDecoder(r.Body).Decode(decoded)
				if !assert.NoError(t, err) {
					t.Fail()
					return
				}
				if !assert.Equal(t, tt.path, r.URL.Path) {
					t.Fail()
					return
				}
				if !assert.Equal(t, tt.data, decoded) {
					t.Fail()
					return
				}
			}))
			defer s.Close()

			(&HTTP{}).Write(&process.M{
				Header: packets.PacketHeader{
					PacketID: tt.packetID,
				},
				Pack: tt.data,
			}, fmt.Sprintf("%s/{{packetID}}", s.URL))

			<-sig
		})
	}
}
