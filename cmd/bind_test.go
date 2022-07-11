package cmd

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

// var ()

// func Test_process_writePacketToHTTP(t *testing.T) {
// 	type args struct {
// 		pack interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		p    process
// 		args args
// 	}{
// 		{
// 			name: "should send json request",
// 			p:    make(process),
// 			args: args{
// 				pack: packets.FinalClassificationData{},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sig := make(chan int)
// 			serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 				var decoded packets.PacketMotionData
// 				err := json.NewDecoder(r.Body).Decode(&decoded)
// 				assert.NoError(t, err)
// 				assert.Equal(t, motionData, decoded)
// 				assert.Equal(t, "/0", r.URL.Path)
// 				close(sig)
// 			}))
// 			defer serv.Close()
// 			tt.p.writePacketToHTTP(struct {
// 				header packets.PacketHeader
// 				raw    interface{}
// 			}{
// 				header: motionData.Header,
// 				raw:    &motionData,
// 			}, fmt.Sprintf("%s/{{packetID}}", serv.URL))
// 			<-sig
// 		})
// 	}
// }

// func Test_process_readPacketsFromConn(t *testing.T) {
// 	stream := make(process, 2)

// 	clientCh := make(chan net.Conn)
// 	go func() {
// 		defer close(clientCh)
// 		binary.Write(<-clientCh, binary.LittleEndian, &motionData)
// 	}()

// 	assert.NoError(t, udptest.NewConn(func(clientConn, serverConn net.Conn) error {
// 		clientCh <- clientConn
// 		stream.readPacketsFromConn(serverConn.(*net.UDPConn), []uint{uint(constants.PacketMotion)})
// 		return nil
// 	}, 2*time.Second))
// 	assert.Equal(t, &motionData, (<-stream).raw)
// }

func Test_reader_read(t *testing.T) {

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
		createTestCase("PacketMotionData21Bytes", mocks.PacketMotionData21Bytes, &mocks.PacketMotionData21),
		createTestCase("PacketParticipantsData21Bytes", mocks.PacketParticipantsData21Bytes, &mocks.PacketParticipantsData21),
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
			r := reader(stream)
			if !assert.NoError(t, udptest.NewConn(func(clientConn, serverConn net.Conn) error {
				clientCh <- clientConn
				(&r).read(ctx, serverConn.(*net.UDPConn), []uint{})
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
