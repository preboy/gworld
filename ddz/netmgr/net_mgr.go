package netmgr

import (
	"net"
	"sync"

	"gworld/core/tcp"
	"gworld/ddz/loop"
	"gworld/ddz/player"
)

var (
	server *tcp.TcpServer
)

var (
	seq      uint32 = 1
	sessions        = map[uint32]*session{}
	lock            = sync.Mutex{}
)

type Session struct {
	Id     uint32
	socket *tcp.Socket
	player *player.Player
}

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		update_chunks()
	})
}

// ----------------------------------------------------------------------------
// export

func Init() {
	server = tcp.NewTcpServer()
	server.Start("0.0.0.0:12345", func(conn *net.TCPConn) {

		sess := NewSession()
		sock := tcp.NewSocket(conn, sess)

		sess.SetSocket(sock)
		sock.Start()
	})
}

func Release() {
	if server != nil {
		server.Stop()
	}

	_lock.Lock()
	defer _lock.Unlock()

	for _, s := range sessions {
		s.Disconnect()
	}
}
