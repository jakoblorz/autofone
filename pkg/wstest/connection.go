package wstest

import (
	"context"
	"net/http/httptest"
	"strings"
	"sync"

	"golang.org/x/net/websocket"
)

func NewConn(fn func(clientConn *websocket.Conn, serverConn *websocket.Conn) error) error {
	serverCh := make(chan *websocket.Conn, 1)
	defer close(serverCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var url string
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		s := httptest.NewServer(websocket.Handler(func(serverConn *websocket.Conn) {
			serverCh <- serverConn
			<-ctx.Done()
		}))
		defer s.Close()
		url = s.URL
		wg.Done()
		<-ctx.Done()
	}()

	wg.Wait()

	clientConn, err := websocket.Dial(strings.ReplaceAll(url, "http", "ws"), "", url)
	if err != nil {
		return err
	}
	defer clientConn.Close()

	return fn(clientConn, <-serverCh)
}
