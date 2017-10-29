package clients

import (
	"net"
)

import (
	"fmt"
)

var _listener *net.TCPListener

func accept() {
	for {
		conn, err := _listener.AcceptTCP()
		if err != nil {
			fmt.Println("error accept")
			break
		}

		OnConnected(conn)
	}
}

func StartListen() bool {
	addr, err := net.ResolveTCPAddr("tcp", ":4040")
	if err != nil {
		return false
	}

	_listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return false
	}

	go accept()

	return true
}

func EndListen() {
	if _listener != nil {
		_listener.Close()
	}
}
