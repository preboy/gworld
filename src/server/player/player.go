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
	server_id   uint32 // 当前Player所运行的server_id
	last_update int64
	run         bool
	w           *sync.WaitGroup
	tf          []func()
}

func NewPlayer() *Player {
	plr := &Player{
		q_packets: make(chan *tcp.Packet, 0x100),
		w:         &sync.WaitGroup{},
	}

	plr.init()
	return plr
}

// --------------------------------------------------------

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

func (self *Player) Stop() {
	self.run = false
	self.w.Wait()
}

func (self *Player) IsRun() bool {
	return self.run
}

// --------------------------------------------------------
// private function

func (self *Player) update() (busy bool) {
	if self.evtMgr.Update() {
		busy = true
	}

	if self.timerMgr.Update() {
		busy = true
	}

	self.last_update = time.Now().UnixNano() / (1000 * 1000)
	return
}

func (self *Player) init() {
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

	for _, fn := range self.tf {
		fn()
	}

	self.tf = self.tf[:0]
}

// --------------------------------------------------------
// public function

func (self *Player) GetSid() int {
	return int(self.sid)
}
