package loop

// the main loop in game

import (
	"fmt"
	"time"

	"core/event"
	"core/schedule"
	"core/thread"
	"core/timer"
	//	"game/modules/act"
)

const (
	_LOOP_NAME = "main_loop"
)

var (
	_thread *thread.Thread
	_loop   *Loop
)

type Loop struct {
	last     int64
	evtMgr   *event.EventMgr
	timerMgr *timer.TimerMgr
}

func (self *Loop) Start() {
	if _thread != nil {
		return
	}

	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)

	schedule.Register(_LOOP_NAME, self)

	_thread = thread.NewThread(func() {
		self.update()
	}, 1000)

	_thread.Go()
}

func (self *Loop) Stop() {
	schedule.UnRegister(_LOOP_NAME)
	if _thread != nil {
		_thread.Stop()
	}

	self.onstop()
}

// ============================================================================
//  IScuedule

func (self *Loop) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Loop) OnEvent(evt *event.Event) {
	switch evt.Id {
	case event.EVT_SCHED_HOUR:
		self.on_new_hour()
	case event.EVT_SCHED_DAY:
		self.on_new_day()
	case event.EVT_SCHED_WEEK:
		self.on_new_week()
	case event.EVT_SCHED_MONTH:
		self.on_new_month()
	case event.EVT_SCHED_YEAR:
		self.on_new_year()
	default:
		fmt.Println("Loop.OnEvent:", evt)
	}
}

func (self *Loop) OnTimer(id uint64) {
	fmt.Println("Loop.OnTimer:", id)
}

func (self *Loop) update() {
	self.evtMgr.Update()
	self.timerMgr.Update()

	sec := time.Now().Unix()
	if self.last != sec {
		self.last = sec
		// todo act.LoopUpdate()
	}
}

// ============================================================================
// public

func (self *Loop) FireEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Loop) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Loop) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

// ============================================================================

func GetLoop() *Loop {
	return _loop
}

func Start() {
	if _loop != nil {
		return
	}

	_loop = &Loop{}
	_loop.Start()
}

func Stop() {
	_loop.Stop()
}
