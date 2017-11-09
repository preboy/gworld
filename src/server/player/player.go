package player

import (
	"time"
)

import (
	"core/event"
	"core/tcp"
)

type Player struct {
	pid         uint64
	aid         uint32
	name        string
	acct        string
	data        uint32
	s           ISession
	evtMgr      *event.EventMgr
	q_packets   chan *tcp.Packet
	last_update int64
	run         bool
	//	quit        chan bool
}

func NewPlayer() *Player {
	plr := &Player{
		q_packets: make(chan *tcp.Packet, 0x100),
	}
	plr.init()
	return plr
}

// ----------------- player evnet -----------------

func (self *Player) Loop() {
	if self.run {
		return
	}
	self.run = true
	defer func() {
		self.run = false
	}()

	for {
		busy := self.dispatch_packet()
		if b := self.update(); b {
			busy = true
		}
		if !busy {
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (self *Player) update() bool {
	now := time.Now().Unix()
	if now-100 >= self.last_update {
		self.last_update = now
		self.on_update()
		return true
	}
	return false
}

func (self *Player) on_update() {
	if self.evtMgr != nil {
		self.evtMgr.Update()
	}
}

func (self *Player) init() {
	// event mgr
	self.evtMgr = event.NewEventMgr(self)

}
