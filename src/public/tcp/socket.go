package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

type IPlayerSocket interface {
	OnRecvPacket(packet *Packet)
}

type Socket struct {
	conn      *net.TCPConn
	w         *sync.WaitGroup
	plr       IPlayerSocket
	q_packets chan *Packet
	on_open   func(*Socket)
	on_closed func(*Socket)
}

func NewSocket(conn *net.TCPConn) *Socket {
	return &Socket{
		conn:      conn,
		w:         &sync.WaitGroup{},
		q_packets: make(chan *Packet),
	}
}

func (self *Socket) SetPlayer(plr IPlayerSocket) {
	self.plr = plr
}

func (self *Socket) Start(on_open, on_closed func(*Socket)) {
	self.on_open = on_open
	self.on_closed = on_closed
	self.on_open(self)
	go self.rt_recv()
	go self.rt_send()
}

func (self *Socket) Stop() {
	self.on_closed(self)
	self.conn.Close()
	self.w.Wait()
}

// running in player's main routine
func (self *Socket) DispatchPacket() bool {
	dispose := false
	for {
		select {
		case packet := <-self.q_packets:
			if self.plr != nil {
				self.plr.OnRecvPacket(packet)
			} else {
				fmt.Println("no plr")
			}
			dispose = true
		default:
			break
		}
	}
	return dispose
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

		// 将数据原样发过去
		// self.conn.Write(head[:])
		// self.conn.Write(body[0:size])

		self.q_packets <- NewPacket(code, body)
		// if plr {
		// 	plr.OnRecvPacket(packet)
		// }
	}

	fmt.Println("socket rt_recv end", self)
}

func (self *Socket) rt_send() {
}

// 发送数据可以另外弄一个routine
func (self *Socket) Send(data []byte) {
	self.conn.Write(data)
}
