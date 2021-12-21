package UoServer

import (
	"fmt"
	"net"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"
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
	
	SERVER_NAME = cfg.Section("server").Key("name").String()
	SERVER_IP = cfg.Section("server").Key("ip").String()
	SERVER_PORT = cfg.Section("server").Key("port").String()
	SOCKET_TIMEOUT = cfg.Section("server").Key("socketTimout").MustInt64()

	fmt.Printf("Server '%v' loading... \n", SERVER_NAME)

	listener, err := net.Listen(CONN_TYPE, SERVER_IP+":"+SERVER_PORT)
	checkError(err, "Fail to create socket server: ")
	fmt.Println("Listening on port " + SERVER_PORT)

	defer listener.Close()

	for server.status != STATUS_STOP {
		conn, err := listener.Accept()
		checkError(err, "Error on connection: ")
		go MonitorConnections(conn)
		
	}
}

func (server *Server) SetStatus(st Status) {
	server.status = st
}

func (server *Server) AddCLient(socket *Socket) {
	u, err := uuid.NewV4()
	checkError(err, "Error while creating uuid to client")

	client := NewClient(socket, u) //Client{}
	
	CLIENTS = append(CLIENTS, client)
	ONLINE_CLIENTS += 1
	fmt.Printf("New Client connected - %v", client.remote_address)
}

func (server *Server) DisconnectClient(id *uuid.UUID) {
	client := GetClientByID(id)

	client.Disconnect()
	RemoveClientByID(client.id)
}

func (server *Server) RemoveClient(id *uuid.UUID) {	
	RemoveClientByID(id)
	ONLINE_CLIENTS -= 1
}


func MonitorConnections(conn net.Conn) {
	if conn != nil {
		socket := NewSocket(conn)
		SERVER.AddCLient(socket)
	}

	millitime := time.Now().UnixMilli()
	for i := 0; i < ONLINE_CLIENTS; i++ {
		client := CLIENTS[i]

		socket := client.socket

		if (millitime - socket.last_input) > SOCKET_TIMEOUT {
			SERVER.DisconnectClient(client.id)
			continue
		} else if client.status != CLIENT_STATUS_CONNECTED {
			continue
		}

		buffer := make([]byte, 128)
		readLength, err := socket.connection.Read(buffer)

		if err != nil {
			fmt.Println("Connection Error: ", err)
		} else if readLength == 0 {
			// TODO: Could add a idle status change to client when its AFK
			continue
		}
	
		fmt.Println("Message Received: " + string(buffer[:len(buffer) - 1]))
		// Packethandling logic
	}
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Println(message, err.Error())
		os.Exit(1)
	}
}
