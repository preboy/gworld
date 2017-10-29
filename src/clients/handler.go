package clients

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// opcode

var CS_PING uint16 = 0x1
var SC_PING uint16 = 0x2

type HANDLER = func(packet *Packet)

var msg_handler [0x10000]HANDLER

func RegisterHandler(opcode uint16, f HANDLER) {
	msg_handler[opcode] = f
}

func handle(packet *Packet) {
	handler := msg_handler[packet.code]
	if handler != nil {
		handler(packet)
	}
}

type Student struct {
	name string
	age  uint32
	man  bool
}

func msg_ping(packet *Packet) {
	var s Student
	buf := bytes.NewBuffer(packet.data)
	dec := gob.NewDecoder(buf)
	dec.Decode(&s)
	fmt.Println(s)
}

func init() {
	RegisterHandler(CS_PING, msg_ping)
}
