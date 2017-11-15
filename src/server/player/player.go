package player

import (
	"time"
)

import (
	"core/event"
	"core/tcp"
	"core/timer"
)

type Player struct {
	pid         uint64
	sid         uint32
	name        string
	acct        string
	data        uint32
	s           ISession
	evtMgr      *event.EventMgr
	timerMgr    *timer.TimerMgr
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

// -------------- private function --------------
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
	self.evtMgr.Update()
	self.timerMgr.Update()
}

func (self *Player) init() {
	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)
}

// -------------- public function --------------
func (self *Player) GetSid() int {
	return int(self.sid)
}
