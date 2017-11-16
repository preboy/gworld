package tcp

import (
	"net"
	"sync/atomic"
)

var (
	_idseq uint32
)

func gen_id() uint32 {
	return atomic.AddUint32(&_idseq, 1)
}

func Connect(host string) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		return nil, err
	}
	return net.DialTCP("tcp", nil, addr)
}

func AsyncConnect(host string, cb func(*net.TCPConn, uint32)) uint32 {
	id := gen_id()
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
