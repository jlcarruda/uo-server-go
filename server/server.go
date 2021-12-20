package UoServer

import (
	"fmt"
	"net"
	"os"

	"gopkg.in/ini.v1"
)

type Status int64

const (
	STATUS_CREATED Status = iota
	STATUS_START
	STATUS_STOP
	STATUS_FATAL
	STATUS_FILE_READING
	STATUS_FILE_READ_FAIL
	STATUS_FILE_READED
	STATUS_FILE_READ_IGNORE
	STATUS_DATABASE_CONNECTING
	STATUS_DATABASE_CONNECTED
	STATUS_DATABASE_CONNECTION_FAILED
	STATUS_UNHANDLED
	STATUS_UNKNOWN
	STATUS_LISTENING
	STATUS_RUNNING

	CONN_TYPE = "tcp"
)

type Server struct {
	status Status
}

func NewServer(testMode bool) *Server {
	if SERVER == nil {
		SERVER = &Server{STATUS_CREATED} 
	}
	return SERVER
}

func (server *Server) Start() {
	server.status = STATUS_START

	cfg, err := ini.Load("./config.ini")
	checkError(err, "Fail to load config file: %v")
	
	name := cfg.Section("server").Key("name").String()
	host := cfg.Section("server").Key("ip").String()
	port := cfg.Section("server").Key("port").String()

	fmt.Printf("Server '%v' loading... \n", name)

	listener, err := net.Listen(CONN_TYPE, host+":"+port)
	checkError(err, "Fail to create socket server: ")
	fmt.Println("Listening on port " + port)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkError(err, "Packet Acception Error: ")
		go HandleConnection(conn)
	}
}

func (server *Server) SetStatus(st Status) {
	server.status = st
}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	
	if err != nil {
		fmt.Println("Connection Error: ", err)
	}

	fmt.Println("Message Received: " + string(buf))

	conn.Close()
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message, err.Error())
		os.Exit(1)
	}
}
