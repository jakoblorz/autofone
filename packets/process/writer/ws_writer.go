package writer

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/jakoblorz/autofone/packets/process"
	"github.com/jakoblorz/autofone/pkg/log"
	"golang.org/x/net/websocket"
)

type Websocket struct {
	*process.P
	*ReceiverRegistry
}

func (s *Websocket) Write(raw interface{}) {
	failedHandles := make([]string, 0)

	s.RLock()
	for handle, conn := range s.conns {
		err := websocket.JSON.Send(conn, raw)
		if err != nil {
			log.Printf("%+v", err)
			failedHandles = append(failedHandles, handle)
		}
	}
	s.RUnlock()

	for _, handle := range failedHandles {
		s.unregisterConn(handle)
	}
}

type ReceiverRegistry struct {
	*sync.RWMutex
	rand.Source
	conns map[string]*websocket.Conn
}

func NewReceiverRegistry() ReceiverRegistry {
	return ReceiverRegistry{
		RWMutex: new(sync.RWMutex),
		Source:  rand.NewSource(time.Now().UnixNano()),
		conns:   make(map[string]*websocket.Conn),
	}
}

func (s *ReceiverRegistry) registerConn(conn *websocket.Conn) string {
	handle := fmt.Sprintf("%s-%f", conn.RemoteAddr(), float64(time.Now().UnixNano())*rand.New(s.Source).Float64())
	s.Lock()
	s.conns[handle] = conn
	s.Unlock()
	return handle
}

func (s *ReceiverRegistry) unregisterConn(handle string) error {
	s.Lock()
	conn, ok := s.conns[handle]
	if !ok {
		s.Unlock()
		return nil
	}
	delete(s.conns, handle)
	s.Unlock()
	return conn.Close()
}

func (s *ReceiverRegistry) Handle(ws *websocket.Conn) {
	handle := s.registerConn(ws)
	defer s.unregisterConn(handle)
	for {
		// discard all messages
		_, err := io.Copy(ioutil.Discard, ws)
		if err != nil {
			log.Printf("%+v", err)
			break
		}
	}
}
