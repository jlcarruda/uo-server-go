package UoServer

import (
	"net"
	"time"
)

const (
	SOCKET_STATUS_CONNECTED int64 = iota
	SOCKET_STATUS_PAUSED
	SOCKET_STATUS_AFK
	SOCKET_STATUS_TIMEDOUT
	SOCKET_STATUS_DISCONNECTED
)

type Socket struct {
	connection net.Conn
	last_input int64
}

func NewSocket(conn net.Conn) *Socket {
	return &Socket{conn, time.Now().UnixMilli()}
}

func (socket *Socket) CloseConnection() {
	socket.connection.Close()
}
