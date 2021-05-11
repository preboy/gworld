package netmgr

import (
	"gworld/core/tcp"
	"net"
	"sync"
)

type session_mgr_t struct {
	dispatcher func(sess *session, packet *tcp.Packet)
	sessions   map[uint32]*session
	server     *tcp.TcpServer
	lock       *sync.Mutex
}

// ----------------------------------------------------------------------------

func (self *session_mgr_t) Init(addr string) {
	self.server.Start(addr, func(conn *net.TCPConn) {
		sess := new_session()
		sock := tcp.NewSocket(conn, sess)
		sess.SetSocket(sock)
		sock.Start()
	})
}

func (self *session_mgr_t) Release() {
	self.lock.Lock()
	defer self.lock.Unlock()

	for _, s := range self.sessions {
		s.Disconnect()
	}

	self.sessions = nil
}

func (self *session_mgr_t) AddSession(sess *session) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.sessions[sess.Id] = sess
}

func (self *session_mgr_t) DelSession(sess *session) {
	self.lock.Lock()
	defer self.lock.Unlock()

	sess.SetMgr(nil)
	delete(self.sessions, sess.Id)
}

func (self *session_mgr_t) OnRecvPacket(sess *session, packet *tcp.Packet) {
	self.dispatcher(sess, packet)
}

func (self *session_mgr_t) SetDispatcher(fn func(*session, *tcp.Packet)) {
	self.dispatcher = fn
}
