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
// 角色线程主体，与玩家网络连接生命周期一致
// 在循环之前需要先执行<时间差操作>
func (self *Player) Go() {
	if self.run {
		return
	}

	go func() {
		self.run = true
		self.w.Add(1)
		_plrs_live[self.data.Acct] = self

		defer func() {
			self.w.Done()
		}()

		// 登录处理
		if self.data.LoginTimes == 0 {
			self.on_first_login()
		}
		self.data.LoginTimes++
		self.on_login()

		self.pursue()

		for self.run {
			busy := self.dispatch_packet()
			if b := self.update(); b {
				busy = true
			}
			if !busy {
				time.Sleep(20 * time.Millisecond)
			}
		}

		self.on_logout()
		self.Save()
		delete(_plrs_live, self.data.Acct)

		println("player.Go exited:", self.data.Name)
	}()

}

func (self *Player) Stop() {
	self.run = false
	self.w.Wait()
}

func (self *Player) IsRun() bool {
	return self.run
}

// -------------- private function --------------
func (self *Player) update() bool {
	now := time.Now().UnixNano() / (1000 * 1000)
	if now >= self.last_update+200 {
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
