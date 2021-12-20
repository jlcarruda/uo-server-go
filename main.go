package main

import (
	UoServer "github.com/jlcarruda/uo-server/server"
)


func main() {
	server := UoServer.NewServer(false)

	server.Start()
}
