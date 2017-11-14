package net_mgr

import (
	"client/session"
	"core/tcp"
)

var (
	socket *tcp.Socket
)

// 新建一个对象，并开启一个route
func on_client_open(conn *tcp.Socket) {

}

func on_client_closed(conn *tcp.Socket) {

}

func Start() {
	conn := tcp.Connect(":4040")
	if conn != nil {
		s := session.NewSession()
		socket = tcp.NewSocket(conn, s)
		s.SetSocket(socket)
		socket.Start(on_client_open, on_client_closed)
		s.Go()
	}
}

func Stop() {
	if socket != nil {
		socket.Stop()
	}
}
