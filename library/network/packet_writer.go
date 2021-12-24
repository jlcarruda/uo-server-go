package Network

import (
	"fmt"
	"sync"
)

type PacketWriter struct {
	capacity int
	stream chan []byte
}

type PacketWriterPool struct {
	lock sync.Mutex
	pool []*PacketWriter
}

// PacketWriterPool ===================================

func (pl *PacketWriterPool) Lock() {
	pl.lock.Lock()
}

func (pl *PacketWriterPool) Unlock() {
	pl.lock.Unlock()
}

func (pl *PacketWriterPool) Pop() *PacketWriter{
	if len(pl.pool) == 0 {
		return nil
	}
	
	pl.Lock()
	pw := pl.pool[len(pl.pool)-1]
	pl.pool = pl.pool[0:len(pl.pool)-2]
	pl.Unlock()

	return pw
}

var pool = PacketWriterPool{pool: []*PacketWriter{}}

// PacketWriter ==========================================
func CreateInstance(capacity int) *PacketWriter {
	pool.Lock()
	var pw = pool.Pop()

	if pw != nil {
		pw.capacity = capacity
		pw.stream = make(chan []byte)
	}

	pool.Unlock()
	if pw == nil {
		pw = &PacketWriter{capacity: capacity, stream: make(chan []byte)}
	}
	fmt.Printf("PacketWriter instance created. There is currently %v writers on pool", len(pool.pool))
	return pw
}

func (pw *PacketWriter) WriteBool(value bool) {
	buffer := make([]byte, 1)
	intValue := 0
	if value {
		intValue = 1
	}	

	pw.stream <- append(buffer, byte(intValue))
}

func (pw *PacketWriter) WriteOneByte(b byte) {
	buffer := make([]byte, 1)
	pw.stream <- append(buffer, b)
}

func (pw *PacketWriter) WriteUint(value uint) {
	buffer := make([]byte, 1)
	pw.stream <- append(buffer, byte(value))
}

func (pw *PacketWriter) WriteInt(value int) {
	buffer := make([]byte, 1)
	pw.stream <- append(buffer, byte(value))
}

func (pw *PacketWriter) WriteSequence(bytes []byte) {
	pw.stream <- bytes
}

