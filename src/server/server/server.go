package server

import (
	"fmt"
)

import (
	"core/event"
	"core/schedule"
	"core/thread"
	"core/timer"
)

const (
	_SERVER_NAME = "main_thread"
)

var (
	_thread *thread.Thread
	_server *Server
)

type Server struct {
	evtMgr   *event.EventMgr
	timerMgr *timer.TimerMgr
}

func NewServer() *Server {
	if _server == nil {
		_server = &Server{}
	}
	return _server
}

func (self *Server) Start() {
	if _thread != nil {
		return
	}

	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)

	schedule.Register(_SERVER_NAME, self)

	_thread = thread.NewThread(game_update, 100)
	_thread.Go()

}

func (self *Server) Stop() {
	schedule.UnRegister(_SERVER_NAME)
	if _thread != nil {
		_thread.Stop()
	}
}

// ----------------- impl for interface

// ----------------- IScuedule -----------------
func (self *Server) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Server) OnEvent(evt *event.Event) {
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
		fmt.Println("Server.OnEvent:", evt)
	}
}

func (self *Server) OnTimer(id uint64) {
	fmt.Println("Server.OnTimer:", id)
}

// ----------------- public -----------------

func (self *Server) FireEvent(evt *event.Event) {
	self.OnEvent(evt)
}

func (self *Server) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Server) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

// ----------------- private -----------------
func game_update() {
	if _server == nil {
		return
	}

	_server.evtMgr.Update()
	_server.timerMgr.Update()
}
