package player

import (
	"time"
)

import (
	"public/event"
	"public/tcp"
)

type Player struct {
	pid         uint32
	aid         uint32
	name        string
	data        uint32
	socket      *tcp.Socket
	evtMgr      *event.EventMgr
	last_update int64
	//	quit        chan bool
}

func NewPlayer() {
	plr = &Player{}
	plr.init()
	return plr
}

// ----------------- player evnet -----------------

func (self *Player) Loop() {
	for {
		busy := false
		if self.socket {
			if b := self.socket.DispatchPacket(); b {
				busy = true
			}
		}
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
		self.evtMgr.Loop()
	}
}

func (self *Player) init() {
	// event mgr
	self.evtMgr = event.NewEventMgr(self)

}
