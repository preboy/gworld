package net_mgr

import (
	"net"
	"sync/atomic"
)

import (
	"core/log"
	"core/tcp"
	"server/game"
	"server/session"
)

var (
	counts_of_client uint32
	server           *tcp.TcpServer
)

// 新建一个对象，并开启一个route
func on_client_connected(conn *net.TCPConn) {
	s := session.NewSession()
	socket := tcp.NewSocket(conn, s)
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
	addr := game.GetServerConfig().Listen_addr
	server = tcp.NewTcpServer()
	server.Start(addr, on_client_connected)
	log.GetLogger().Info("server listen on: %s", addr)
}

func Stop() {
	if server != nil {
		server.Stop()
	}
}
