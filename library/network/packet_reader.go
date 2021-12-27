package Network

type PacketReader struct {
	data []byte
	size int
	index int
}

func CreatePacketReader(data []byte, size int, fixedSize bool) *PacketReader {
	index := 3
	if fixedSize {
		index = 1
	}
	return &PacketReader{
		data,
		size,
		index,
	}
}

// TODO: Trace uses some other dependencies. Going to implement later. Refer to NetState.cs 57

