package wstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
)

type TestMsg struct {
	Value uint16
}

func TestNewConn(t *testing.T) {
	clientCh := make(chan *websocket.Conn)
	go func() {
		defer close(clientCh)
		websocket.JSON.Send(<-clientCh, &TestMsg{12})
	}()

	var received TestMsg
	assert.NoError(t, NewConn(func(clientConn, serverConn *websocket.Conn) error {
		clientCh <- clientConn
		return websocket.JSON.Receive(serverConn, &received)
	}))
	assert.Equal(t, uint16(12), received.Value)
}
