package Network

import (
	uuid "github.com/nu7hatch/gouuid"
)

type PacketStatus byte

const (
	PACKET_INACTIVE PacketStatus = iota
	PACKET_STATIC
	PACKET_ACQUIRED
	PACKET_ACCESSED
	PACKET_BUFFERED
	PACKET_WARNED
)

type Packet struct {
	status PacketStatus
	id *uuid.UUID
	length int
}
