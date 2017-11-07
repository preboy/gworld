package player

import (
	"time"
)

import (
	"core/event"
	"core/tcp"
)

type Player struct {
	pid         uint32
	aid         uint32
	name        string
	data        uint32
	socket      *tcp.Socket
	evtMgr      *event.EventMgr
	q_packets   chan *Packet
	last_update int64
	//	quit        chan bool
}

func NewPlayer() {
	plr = &Player{
		q_packets: make(chan *Packet, 0x100),
	}
	plr.init()
	return plr
}

// ----------------- player evnet -----------------

func (self *Player) Loop() {
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
	if self.evtMgr {
		self.evtMgr.Update()
	}
}

func (self *Player) init() {
	// event mgr
	self.evtMgr = event.NewEventMgr(self)

}
