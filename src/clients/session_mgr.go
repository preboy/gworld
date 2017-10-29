package clients

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"
)

var counts_of_connection uint32

type Packet struct {
	code uint16
	data []byte
}

func dispatch_packet(packet *Packet) {
	fmt.Println("new data:", packet.code, len(packet.data))
	/*
		for i := 0; i < len(data); i++ {
			fmt.Printf("%x ", data[i])
		}
		fmt.Println()
	*/

}

func session_loop(conn *net.TCPConn) {
	defer conn.Close()

	var plr *Player = CreatePlayer("name")

	for {
		head := make([]byte, 4)
		var l int = 0
		for l < 4 {
			len, err := conn.Read(head[l:4])
			if err != nil {
				fmt.Println("read err:", err)
				return
			}
			l += len
		}
		buff := bytes.NewReader(head)

		var size uint16
		var code uint16
		binary.Read(buff, binary.LittleEndian, &size)
		binary.Read(buff, binary.LittleEndian, &code)

		body := make([]byte, size)
		l = 0
		for uint16(l) < size {
			len, err := conn.Read(body[l:size])
			if err != nil {
				fmt.Println("read err:", err)
				return
			}
			l += len
		}

		// 将数据原样发过去
		conn.Write(head[:])
		conn.Write(body[0:size])

		packet := new(Packet)
		packet.code = code
		packet.data = body[0:size]
		dispatch_packet(packet)
		body = nil
		plr.ch_packet <- *packet
	}

	fmt.Println("socketector terminated", conn)
}

// 新建一个对象，并开启一个route
func OnConnected(conn *net.TCPConn) {
	atomic.AddUint32(&counts_of_connection, 1)
	go session_loop(conn)
}

func OnDisconnected() {
	atomic.AddUint32(&counts_of_connection, ^uint32(0))
}
