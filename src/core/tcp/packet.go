package tcp

type Packet struct {
	opcode uint16
	data   []byte
}

func NewPacket(opcode uint16, data []byte) *Packet {
	return &Packet{
		opcode: opcode,
		data:   data,
	}
}
