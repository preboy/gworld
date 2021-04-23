package netmgr

import (
	"net"
	"sync"

	"gworld/core/tcp"
	"gworld/game/player"
	"gworld/game/session"
)

var (
	server *tcp.TcpServer
)

var (
	seq          uint32 = 1
	all_sessions        = map[uint32]*Session{}
	lock                = sync.Mutex{}
)

type Session struct {
	Id     uint32
	socket *tcp.Socket
	player *player.Player
}

// ----------------------------------------------------------------------------
// export

func Init() {
	server = tcp.NewTcpServer()
	server.Start("0.0.0.0:12345", func(conn *net.TCPConn) {

		sess := session.NewSession()
		sock := tcp.NewSocket(conn, sess)

		sess.SetSocket(sock)
		sock.Start()
	})
}

func Release() {
	if server != nil {
		server.Stop()
	}

	session.Stop()
}
