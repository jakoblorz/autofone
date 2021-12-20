package cmd

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/jakoblorz/f1-metrics-transformer/pkg/udptest"
	"github.com/stretchr/testify/assert"
)

var ()

func Test_process_writePacketToHTTP(t *testing.T) {
	type args struct {
		pack interface{}
	}
	tests := []struct {
		name string
		p    process
		args args
	}{
		{
			name: "should send json request",
			p:    make(process),
			args: args{
				pack: packets.FinalClassificationData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := make(chan int)
			serv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				var decoded packets.PacketMotionData
				err := json.NewDecoder(r.Body).Decode(&decoded)
				assert.NoError(t, err)
				assert.Equal(t, motionData, decoded)
				close(sig)
			}))
			defer serv.Close()
			tt.p.writePacketToHTTP(&motionData, serv.URL)
			<-sig
		})
	}
}

func Test_process_readPacketsFromConn(t *testing.T) {
	stream := make(process, 2)

	clientCh := make(chan net.Conn)
	go func() {
		defer close(clientCh)
		binary.Write(<-clientCh, binary.LittleEndian, &motionData)
	}()

	assert.NoError(t, udptest.NewConn(func(clientConn, serverConn net.Conn) error {
		clientCh <- clientConn
		stream.readPacketsFromConn(serverConn.(*net.UDPConn), []uint{uint(env.PacketMotion)})
		return nil
	}, 2*time.Second))
	assert.Equal(t, &motionData, <-stream)
}
