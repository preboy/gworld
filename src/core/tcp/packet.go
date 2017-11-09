package tcp

type Packet struct {
	Opcode uint16
	Data   []byte
}

func NewPacket(opcode uint16, data []byte) *Packet {
	return &Packet{
		Opcode: opcode,
		Data:   data,
	}
}
