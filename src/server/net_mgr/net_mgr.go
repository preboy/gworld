package net_mgr

import (
	"net"
)

import (
	"core/log"
	"core/tcp"
	"server/game"
	"server/session"
)

var (
	server *tcp.TcpServer
)

// 新建一个对象，并开启一个route
func on_client_connected(conn *net.TCPConn) {
	s := session.NewSession()
	socket := tcp.NewSocket(conn, s)
	s.SetSocket(socket)
	socket.Start()
}

func Start() {
	addr := game.GetServerConfig().Listen_addr
	server = tcp.NewTcpServer()
	server.Start(addr, on_client_connected)
	log.Info("server listen on: %s", addr)
}

func Stop() {
	if server != nil {
		server.Stop()
	}
}
