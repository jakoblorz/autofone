package udptest

import (
	"net"
	"time"
)

func NewConn(fn func(clientConn net.Conn, serverConn net.Conn) error, timeout time.Duration) error {
	var (
		readDeadline = time.Now().Add(timeout)
	)

	serverConn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return err
	}
	defer serverConn.Close()

	clientConn, err := net.DialTimeout("udp", serverConn.LocalAddr().String(), timeout)
	if err != nil {
		return err
	}
	defer clientConn.Close()

	_ = serverConn.SetReadDeadline(readDeadline)
	_ = clientConn.SetWriteDeadline(readDeadline)

	return fn(clientConn, serverConn)
}
