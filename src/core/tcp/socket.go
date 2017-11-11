package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

type ISession interface {
	OnRecvPacket(packet *Packet)
}

type Socket struct {
	conn      *net.TCPConn
	w         *sync.WaitGroup
	s         ISession
	fn_open   func(*Socket)
	fn_closed func(*Socket)
}

func NewSocket(conn *net.TCPConn, s ISession) *Socket {
	return &Socket{
		conn: conn,
		s:    s,
		w:    &sync.WaitGroup{},
	}
}

func (self *Socket) Start(oopen, closed func(*Socket)) {
	self.fn_open = oopen
	self.fn_closed = closed

	if self.fn_open != nil {
		self.fn_open(self)
	}

	go self.rt_recv()
	go self.rt_send()
}

func (self *Socket) Stop() {
	if self.fn_closed != nil {
		self.fn_closed(self)
	}

	self.conn.Close()
	self.w.Wait()
}

func (self *Socket) rt_recv() {
	defer func() {
		self.w.Done()
	}()

	self.w.Add(1)

	for {
		head := make([]byte, 4)
		var l int = 0
		for l < 4 {
			len, err := self.conn.Read(head[l:4])
			if err != nil {
				fmt.Println("read err:", err)
				break
			}
			l += len
		}
		buff := bytes.NewReader(head)

		var size uint16
		var code uint16
		binary.Read(buff, binary.LittleEndian, &size)
		binary.Read(buff, binary.LittleEndian, &code)

		l = 0
		body := make([]byte, size)
		for uint16(l) < size {
			len, err := self.conn.Read(body[l:size])
			if err != nil {
				fmt.Println("read err:", err)
				break
			}
			l += len
		}
		self.s.OnRecvPacket(NewPacket(code, body))
	}

	fmt.Println("socket rt_recv end", self)
}

func (self *Socket) rt_send() {
}

// 发送数据可以另外弄一个routine
func (self *Socket) Send(data []byte) {
	self.conn.Write(data)
}
