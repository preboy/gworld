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
	sndq chan *[]byte
	stop bool
}

func NewSocket(conn *net.TCPConn, s ISession) *Socket {
	// conn.SetNoDelay(false)
	return &Socket{
		conn: conn,
		s:    s,
		w:    &sync.WaitGroup{},
		sndq: make(chan *[]byte, 128),
		stop: false,
	}
}

func (self *Socket) Start() {
	self.s.OnOpened()
	go self.rt_recv()
	go self.rt_send()
}

func (self *Socket) Stop() {
	if self.stop {
		return
	}

	self.stop = true
	self.s.OnClosed()

	if self.sndq != nil {
		close(self.sndq)
		self.sndq = nil
	}

	if self.conn != nil {
		self.conn.Close()
		self.conn = nil
	}

	self.w.Wait()
}

func (self *Socket) rt_recv() {
	self.w.Add(1)
	defer func() {
		self.w.Done()
		self.Stop()
	}()

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
}

func (self *Socket) rt_send() {
	self.w.Add(1)
	defer func() {
		self.w.Done()
		self.Stop()
	}()

	for {
		select {
		case buf, ok := <-self.sndq:
			if !ok {
				return
			}

			L := len(self.sndq)
			for L > 0 && len(*buf) < 4096 {
				*buf = append(*buf, *<-self.sndq...)
				L--
			}

			n, err := self.conn.Write(*buf)
			if err != nil {
				fmt.Println("send data failed !")
				return
			}

			if n != len(*buf) {
				fmt.Println("send data unfinished !")
				return
			}

			if L == 0 {
				time.Sleep(time.Duration(100) * time.Millisecond)
			}
		}
	}
}

func (self *Socket) Send(data []byte) {
	defer func() {
		fmt.Println("Send Error !")
		recover()
	}()

	self.sndq <- &data
}
