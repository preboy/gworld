package microsvr

import (
	"sync"
	"time"

	"core/event"
	"core/schedule"
	"core/tcp"
	"core/timer"
	"game/app"
)

type IMod interface {
	GetName() string
	GetOpcodeRange() (uint16, uint16)

	OnStart()
	OnStop()
	OnUpdate()

	OnTimer(id uint64)
	OnEvent(evt *event.Event)
}

type msg struct {
	plr    app.IPlayer
	packet *tcp.Packet
}

type Svr struct {
	w        *sync.WaitGroup
	mod      IMod
	quit     chan bool
	msgc     chan *msg
	evtMgr   *event.EventMgr
	timerMgr *timer.TimerMgr
}

// ----------------------------------------------------------------------------

func (self *Svr) Start() {
	go func() {

		self.w.Add(1)
		defer func() {
			self.w.Done()
		}()

		schedule.Register(self.mod.GetName(), self)
		self.evtMgr = event.NewEventMgr(self)
		self.timerMgr = timer.NewTimerMgr(self)

		self.mod.OnStart()

	EXIT:
		for {
			select {
			case <-self.quit:
				break EXIT
			case msg, ok := <-self.msgc:
				if ok {
					msg.plr.DoPacket(msg.packet)
				} else {
					break EXIT
				}
			default:
				time.Sleep(time.Duration(1000) * time.Millisecond)
			}

			self.mod.OnUpdate()

			// event & timer
			self.evtMgr.Update()
			self.timerMgr.Update()
		}

		self.mod.OnStop()
	}()
}

func (self *Svr) Stop() {
	self.quit <- true
	close(self.msgc)
	self.w.Wait()
}

func (self *Svr) PostPacket(plr app.IPlayer, packet *tcp.Packet) {
	// todo: may be panic
	self.msgc <- &msg{plr, packet}
}

// ----------------------------------------------------------------------------
// public

func (self *Svr) FireEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Svr) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Svr) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

// ----------------------------------------------------------------------------

func (self *Svr) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Svr) OnTimer(id uint64) {
	self.mod.OnTimer(id)
}

func (self *Svr) OnEvent(evt *event.Event) {
	self.mod.OnEvent(evt)
}
