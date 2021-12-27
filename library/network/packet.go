package Network

import (
	"fmt"
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
	id int
	length int
	stream *PacketWriter
}

func CreatePacket(id int, length int) *Packet {
	return &Packet{
		id: id,
		length: length,
	}
}

func (p *Packet) SetStream(stream *PacketWriter) {
	p.stream = stream
	p.stream.WriteInt(p.id)
}

func (p *Packet) ValidatePacketLength() {
	if p.length == 0 {
		p.length = p.stream.capacity
	} else if p.length != p.stream.capacity {

		diff := p.length - p.stream.capacity
		
		signal := "-"
		if diff >= 0 {
			signal = "+"
		}

		fmt.Printf("Packet %v: Bad packet length! (%v%v bytes)", p.id, signal, diff)
	}
}
