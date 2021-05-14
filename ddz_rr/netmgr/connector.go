package netmgr

import (
	"bytes"
	"encoding/binary"
	"net"

	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/comp"

	"github.com/gogo/protobuf/proto"
)

type connector struct {
	id       uint32
	socket   *tcp.Socket
	packet_q chan *tcp.Packet

	events  map[string]func(*connector, []interface{})
	handler func(*connector, *tcp.Packet)
}

// ----------------------------------------------------------------------------
// member

func NewConnector(h func(*connector, *tcp.Packet)) *connector {
	return &connector{
		handler: h,
		events:  map[string]func(*connector, []interface{}){},
	}
}

func (self *connector) Start(addr string) {
	self.packet_q = make(chan *tcp.Packet, 0x100)

	self.id = tcp.AsyncConnect(addr, func(conn *net.TCPConn, err error) {
		if conn == nil {
			self.Fire("error", err.Error())
			return
		}

		self.socket = tcp.NewSocket(conn, self)
		self.socket.Start()
	})
}

func (self *connector) Close() {
	self.socket.Stop()
}

func (self *connector) Update() {
	for {
		select {
		case p := <-self.packet_q:
			self.handler(self, p)
		default:
			return
		}
	}
}

func (self *connector) On(evt string, fn func(*connector, []interface{})) *connector {
	self.events[evt] = fn
	return self
}

func (self *connector) Fire(evt string, args ...interface{}) {
	if fn, ok := self.events[evt]; ok {
		fn(self, args)
	}
}

// ----------------------------------------------------------------------------
// impl for ISession

func (self *connector) OnRecvPacket(packet *tcp.Packet) {
	self.packet_q <- packet
}

func (self *connector) OnOpened() {
	self.Fire("opened")
}

func (self *connector) OnClosed() {
	self.Fire("closed")
}

func (self *connector) SendPacket(opcode uint16, data []byte) {
	if self.socket == nil {
		return
	}

	l := uint16(len(data))
	b := make([]byte, 0, l+2+2)
	buf := bytes.NewBuffer(b)
	binary.Write(buf, binary.LittleEndian, l)
	binary.Write(buf, binary.LittleEndian, opcode)
	binary.Write(buf, binary.LittleEndian, data)
	self.socket.Send(buf.Bytes())
}

func (self *connector) SendMessage(msg comp.IMessage) {
	str := utils.ObjectToString(msg)
	log.Info("rr SEND packet: %v, %v, %v", self.id, msg.GetOP(), str)

	self.SendProtobufMessage(uint16(msg.GetOP()), msg)
}

func (self *connector) SendProtobufMessage(opcode uint16, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("proto.Marshal ERROR: %v %v %v", self.id, opcode, err)
		return
	}

	self.SendPacket(opcode, data)
}
