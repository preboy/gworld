package comp

type ISession interface {
	SendPacket(opcode uint16, data []byte)
}
