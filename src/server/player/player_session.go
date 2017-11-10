package player

import (
	"bytes"
	"encoding/binary"
	"fmt"
)
import (
	"github.com/gogo/protobuf/proto"
)

type ISession interface {
	Send(data []byte)
	SetPlayer(*Player)
	Disconnect()
}

func (self *Player) SetSession(s ISession) {
	self.s = s
}

func (self *Player) Disconnect() {
	self.s.Disconnect()
	self.s = nil
}

func (self *Player) SendPacket(opcode uint16, obj proto.Message) {
	if self.s == nil {
		return
	}

	data, err := proto.Marshal(obj)
	if err == nil {
		l := uint16(len(data))
		b := make([]byte, 0, l+2+2)
		buf := bytes.NewBuffer(b)
		binary.Write(buf, binary.LittleEndian, uint16(len(data)))
		binary.Write(buf, binary.LittleEndian, opcode)
		binary.Write(buf, binary.LittleEndian, data)
		self.s.Send(buf.Bytes())
	} else {
		fmt.Println("SendPacket Error:failed to Marshal obj")
	}

}
