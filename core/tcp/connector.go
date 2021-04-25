package tcp

import (
	"gworld/core/utils"
	"net"
)

func Connect(host string) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		return nil, err
	}
	return net.DialTCP("tcp", nil, addr)
}

func AsyncConnect(host string, cb func(*net.TCPConn, uint32)) uint32 {
	id := utils.SeqU32()
	go func() {
		addr, err := net.ResolveTCPAddr("tcp4", host)
		if err != nil {
			cb(nil, id)
		} else {
			conn, err := net.DialTCP("tcp", nil, addr)
			if err == nil {
				cb(conn, id)
			} else {
				cb(nil, id)
			}
		}
	}()
	return id
}
