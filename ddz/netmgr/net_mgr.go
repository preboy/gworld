package netmgr

import (
	"strconv"
	"sync"

	"gworld/core/tcp"
	"gworld/ddz/comp"
	"gworld/ddz/config"
	"gworld/ddz/loop"
)

var (
	_gambler_mgr *session_mgr_t
	_referee_mgr *session_mgr_t
)

var (
	_gambler_chunks = make(chan *chunk, 0x1000)
	_referee_chunks = make(chan *chunk, 0x1000)
)

type chunk struct {
	s *session
	p *tcp.Packet
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
	// gambler
	_gambler_mgr = &session_mgr_t{
		lock:     &sync.Mutex{},
		server:   tcp.NewTcpServer(),
		sessions: map[uint32]*session{},
	}

	_gambler_mgr.SetDispatcher(func(s *session, p *tcp.Packet) {
		if s.player != nil {
			_gambler_chunks <- &chunk{s, p}
		} else {
			pid := strconv.Itoa(int(s.Id))
			plr := comp.GM.NewGambler(pid)
			s.SetPlayer(plr)

			loop.Post(func() {
				plr.OnLogin()
			})
		}
	})

	_gambler_mgr.Init(config.Get().Addr4Gambler)

	// referee
	_referee_mgr = &session_mgr_t{
		lock:     &sync.Mutex{},
		server:   tcp.NewTcpServer(),
		sessions: map[uint32]*session{},
	}

	_referee_mgr.SetDispatcher(func(s *session, p *tcp.Packet) {

		if s.player == nil {
			pid := strconv.Itoa(int(s.Id))
			plr := comp.RM.NewReferee(pid)
			s.SetPlayer(plr)

			loop.Post(func() {
				plr.OnLogin()
			})
		}

		_referee_chunks <- &chunk{s, p}
	})

	_referee_mgr.Init(config.Get().Addr4Referee)
}

func Release() {
	_gambler_mgr.Release()
	_referee_mgr.Release()
}

// ----------------------------------------------------------------------------
// local

func update_chunks() {
	for {
		select {
		case c := <-_gambler_chunks:
			c.s.player.OnPacket(c.p)
			loop.DoNext()
		case c := <-_referee_chunks:
			c.s.player.OnPacket(c.p)
			loop.DoNext()
		default:
			return
		}
	}
}
