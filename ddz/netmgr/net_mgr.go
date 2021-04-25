package netmgr

import (
	"net"

	"gworld/core/tcp"
	"gworld/ddz/loop"
)

var (
	server *tcp.TcpServer
)

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

		sess := new_session()
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

	for _, s := range _sessions {
		s.Disconnect()
	}
}
