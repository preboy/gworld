package netmgr

import (
	"fmt"
	"net"

	"core/log"
	"core/tcp"
	"game/app"
	"game/session"
)

var (
	server *tcp.TcpServer
)

func Start() {
	addr := fmt.Sprintf("%s:%d", app.GetGameConfig().Host, app.GetGameConfig().Port)
	server = tcp.NewTcpServer()
	server.Start(addr, func() {
		s := session.NewSession()
		socket := tcp.NewSocket(conn, s)
		s.SetSocket(socket)
		socket.Start()
	})
	log.Info("server listen on: %s", addr)
}

func Stop() {
	if server != nil {
		server.Stop()
	}
}
