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

	addr := fmt.Sprintf("0.0.0.0:%d", app.GetGameConfig().Port)
	server = tcp.NewTcpServer()

	server.Start(addr, func(conn *net.TCPConn) {

		sess := session.NewSession()
		sock := tcp.NewSocket(conn, sess)

		sess.SetSocket(sock)
		sock.Start()
	})

	log.Info("server listen on: %s", addr)
}

func Stop() {
	if server != nil {
		server.Stop()
	}

	session.Stop()
}
