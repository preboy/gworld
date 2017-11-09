package net_mgr

import (
	"net"
	"sync/atomic"
)

import (
	"core/tcp"
	"server/session"
)

var (
	counts_of_client uint32
	server           *tcp.TcpServer
)

// 新建一个对象，并开启一个route
func on_client_connected(conn *net.TCPConn) {
	s := session.NewSession()
	socket := tcp.NewSocket(conns, s)
	s.SetSocket(socket)
	socket.Start(on_client_open, on_client_closed)
}

func on_client_open(socket *tcp.Socket) {
	atomic.AddUint32(&counts_of_client, 1)
}

func on_client_closed(socket *tcp.Socket) {
	atomic.AddUint32(&counts_of_client, ^uint32(0))
}

func Start() {
	server = tcp.NewTcpServer()
	server.Start(":4040", on_client_connected)
}

func Stop() {
	if server != nil {
		server.Stop()
	}
}
