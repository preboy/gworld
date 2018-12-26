package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

type ISession interface {
	OnRecvPacket(packet *Packet)
	OnOpened()
	OnClosed()
}

type Socket struct {
	conn *net.TCPConn
	w    *sync.WaitGroup
	s    ISession
}

func NewSocket(conn *net.TCPConn, s ISession) *Socket {
	// conn.SetNoDelay(false)
	return &Socket{
		conn: conn,
		s:    s,
		w:    &sync.WaitGroup{},
	}
}

func (self *Socket) Start() {
	self.s.OnOpened()
	go self.rt_recv()
	go self.rt_send()
}

func (self *Socket) Stop() {
	self.conn.Close()
	self.w.Wait()
}

func (self *Socket) rt_recv() {
	defer func() {
		self.w.Done()
	}()

	self.w.Add(1)
J:
	for {
		self.conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

		head := make([]byte, 4)
		var l int = 0
		for l < 4 {
			n, err := self.conn.Read(head[l:4])
			if err != nil || n == 0 {
				fmt.Println("read err:", err)
				break J
			}
			l += n
		}
		buff := bytes.NewReader(head)

		var size uint16
		var code uint16
		binary.Read(buff, binary.LittleEndian, &size)
		binary.Read(buff, binary.LittleEndian, &code)

		l = 0
		body := make([]byte, size)
		for uint16(l) < size {
			n, err := self.conn.Read(body[l:size])
			if err != nil || n == 0 {
				fmt.Println("read err:", err)
				break J
			}
			l += n
		}
		self.s.OnRecvPacket(NewPacket(code, body))
	}

	self.s.OnClosed()
	self.conn.Close()
}

func (self *Socket) rt_send() {
}

// 发送数据可以另外弄一个routine
func (self *Socket) Send(data []byte) {
	n, err := self.conn.Write(data)
	if err != nil {
		fmt.Println("Write err:", err)
		self.Stop()
	} else if n < len(data) {
		fmt.Println("Write dealy")
		time.Sleep(100 * time.Millisecond)
		self.Send(data[n:])
	}
}
