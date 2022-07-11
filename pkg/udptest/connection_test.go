package udptest

import (
	"encoding/binary"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestMsg struct {
	Value uint16
}

func TestNewConn(t *testing.T) {
	clientCh := make(chan net.Conn)
	go func() {
		defer close(clientCh)
		binary.Write(<-clientCh, binary.LittleEndian, &TestMsg{12})
	}()

	var received TestMsg
	assert.NoError(t, NewConn(func(clientConn, serverConn net.Conn) error {
		clientCh <- clientConn
		return binary.Read(serverConn, binary.LittleEndian, &received)
	}, time.Second))
	assert.Equal(t, uint16(12), received.Value)

}
