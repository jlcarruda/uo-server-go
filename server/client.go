package UoServer

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type ClientStatus int64

const (
	CLIENT_STATUS_CONNECTED ClientStatus = iota
	CLIENT_STATUS_DISCONNECTED
)


type Client struct {
	socket *Socket
	remote_address string
	id *uuid.UUID
	status ClientStatus
}

func NewClient(socket *Socket, id *uuid.UUID) *Client {
	client := Client{}
	client.id = id
	client.socket = socket
	client.remote_address = socket.connection.RemoteAddr().String()
	client.status = CLIENT_STATUS_CONNECTED

	return &client
}


func (client *Client) Send(b []byte) {
	_, err := client.socket.connection.Write(b)
	if err != nil {
		fmt.Println("Error while sending message to client: " + err.Error())
	}
}

func (client *Client) Disconnect() {
	client.Send([]byte("disconnect"))
	client.socket.CloseConnection()
	client.status = CLIENT_STATUS_DISCONNECTED
}
