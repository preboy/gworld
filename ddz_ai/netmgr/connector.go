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

func (c *connector) Start(addr string) {
	c.packet_q = make(chan *tcp.Packet, 0x100)

	c.id = tcp.AsyncConnect(addr, func(conn *net.TCPConn, err error) {
		if conn == nil {
			c.Fire("error", err.Error())
			return
		}

		c.socket = tcp.NewSocket(conn, c)
		c.socket.Start()
	})
}

func (c *connector) Close() {
	c.socket.Stop()
}

func (c *connector) Update() {
	for {
		select {
		case p := <-c.packet_q:
			c.handler(c, p)
		default:
			return
		}
	}
}

func (c *connector) On(evt string, fn func(*connector, []interface{})) *connector {
	c.events[evt] = fn
	return c
}

func (c *connector) Fire(evt string, args ...interface{}) {
	if fn, ok := c.events[evt]; ok {
		fn(c, args)
	}
}

// ----------------------------------------------------------------------------
// impl for ISession

func (c *connector) OnRecvPacket(packet *tcp.Packet) {
	c.packet_q <- packet
}

func (c *connector) OnOpened() {
	c.Fire("opened")
}

func (c *connector) OnClosed() {
	c.Fire("closed")
}

func (c *connector) SendPacket(opcode uint16, data []byte) {
	if c.socket == nil {
		return
	}

	l := uint16(len(data))
	b := make([]byte, 0, l+2+2)
	buf := bytes.NewBuffer(b)
	binary.Write(buf, binary.LittleEndian, l)
	binary.Write(buf, binary.LittleEndian, opcode)
	binary.Write(buf, binary.LittleEndian, data)
	c.socket.Send(buf.Bytes())
}

func (c *connector) SendMessage(msg comp.IMessage) {
	str := utils.ObjectToString(msg)
	log.Info("SEND packet: %v, %v, %v", c.id, msg.GetOP(), str)

	c.SendProtobufMessage(uint16(msg.GetOP()), msg)
}

func (c *connector) SendProtobufMessage(opcode uint16, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("proto.Marshal ERROR: %v %v %v", c.id, opcode, err)
		return
	}

	c.SendPacket(opcode, data)
}
