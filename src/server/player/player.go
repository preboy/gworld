package player

import (
	"time"
)

import (
	"core/event"
	"core/tcp"
	"core/timer"
	"sync"
)

type Player struct {
	sid         uint32
	data        *PlayerData
	s           ISession
	evtMgr      *event.EventMgr
	timerMgr    *timer.TimerMgr
	q_packets   chan *tcp.Packet
	last_update int64
	run         bool
	w           *sync.WaitGroup
}

func NewPlayer() *Player {
	plr := &Player{
		q_packets: make(chan *tcp.Packet, 0x100),
		w:         &sync.WaitGroup{},
	}
	plr.init()
	return plr
}

// ----------------- player evnet -----------------
func (self *Player) Go() {
	if self.run {
		return
	}

	self.run = true
	self.w.Add(1)

	defer func() {
		self.run = false
		self.w.Done()
	}()

	for {
		busy := self.dispatch_packet()
		if b := self.update(); b {
			busy = true
		}
		if !busy {
			if !self.run {
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (self *Player) Stop() {
	self.run = false
	self.w.Wait()
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
