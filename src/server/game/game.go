package game

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
	SERVER_NAME = "Game"
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

func (self *Server) Init() {
	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)

	schedule.Register(SERVER_NAME, self)
}

func (self *Server) Start() {
	if _thread == nil {
		_thread = thread.NewThread(game_update, 100)
		_thread.Go()
	}
}

func (self *Server) Stop() {
	schedule.UnRegister(SERVER_NAME)
	if _thread != nil {
		_thread.Stop()
	}
}

// ----------------- impl for interface

// ----------------- IScuedule -----------------
func (self *Server) EmitEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Server) FireEvent(evt *event.Event) {
	self.OnEvent(evt)
}

func (self *Server) OnEvent(evt *event.Event) {
	fmt.Println("Server.OnEvent:", evt)
}

func (self *Server) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Server) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

func (self *Server) OnTimer(id uint64) {

}

// ----------------- private -----------------
func game_update() {
	if _server == nil {
		return
	}

	_server.evtMgr.Update()
	_server.timerMgr.Update()
}
