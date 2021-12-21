package UoServer

var (
	SERVER *Server
	SERVER_NAME string
	SERVER_PORT string
	SERVER_IP string
	CLIENTS []*Client
	ONLINE_CLIENTS int
	SOCKET_TIMEOUT int64
)
