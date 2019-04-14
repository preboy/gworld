package player

import (
	"sync"
	"time"

	"core/event"
	"core/log"
	"core/tcp"
	"core/timer"
	"core/utils"
)

type Player struct {
	s    ISession
	tf   []func()
	sid  uint32 // slot id
	data *PlayerData
}

func NewPlayer() *Player {
	plr := &Player{}
	return plr
}

// ============================================================================

// player event
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
			if self.update() {
				busy = true
			}
			if !busy && self.idle() {
				time.Sleep(20 * time.Millisecond)
			}
		}

		self.on_logout()
		self.Save()
		_plrs_live[self.data.Acct] = nil

		println("player.Go exited:", self.data.Name)
	}()

}

// ============================================================================
// private function

func (self *Player) Init() {
	self.data.Init(self)
	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)
	self.tf = make([]func(), 0, 4)
}

// checking whether idle or not the loop
func (self *Player) idle() bool {
	if len(self.q_packets) != 0 {
		return false
	}

	if self.evtMgr.Len() != 0 {
		return false
	}

	// timerMgr ignored
	return true
}

func (self *Player) do_next_tick() {
	if len(self.tf) == 0 {
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Error("PANIC on 'do_next_tick':", self.GetId())
			log.Error("STACK TRACE:", utils.Callstack())
		}
	}()

	for _, fn := range self.tf {
		fn()
	}

	self.tf = self.tf[:0]
}

// ============================================================================
// public function

func (self *Player) GetSid() int {
	return int(self.sid)
}
