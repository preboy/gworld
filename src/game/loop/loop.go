package loop

// the main loop in game

import (
	"fmt"
	"time"

	"core/event"
	"core/schedule"
	"core/timer"
)

const (
	_LOOP_NAME = "main_loop"
)

var (
	_loop *Loop
)

// ============================================================================

type Loop struct {
	q        chan bool
	w        *sync.WaitGroup
	last     int64
	talks    chan *talk
	evtMgr   *event.EventMgr
	timerMgr *timer.TimerMgr
}

// ============================================================================

type iSession interface {
	DoPacket(*tcp.Packet)
}

type talk struct {
	s iSession
	p *tcp.Packet
}

func (self *Loop) start() {
	self.q = make(chan bool)
	self.w = &sync.WaitGroup{}
	self.talks = make(chan *talk, 0x1000)
	self.evtMgr = event.NewEventMgr(self)
	self.timerMgr = timer.NewTimerMgr(self)

	go func() {
		self.working()
	}()
}

func (self *Loop) stop() {
	self.q <- true
	self.w.Wait()
}

// ============================================================================

func (self *Loop) working() {
	self.w.Add(1)
	defer func() {
		self.w.Done()
	}()

	self.run = true
	defer func() {
		self.run = false
	}()

	self.on_start()

	var busy bool

	for {

		busy = false

		select {
		case <-self.quit:
			break
		case talk := <-self.talks:
			self.do_talk(talk)
			busy = true
		default:
		}

		if self.evtMgr.Update() {
			busy = true
		}

		if self.timerMgr.Update() {
			busy = true
		}

		if self.do_update() {
			busy = true
		}

		if !busy {
			busy = self.do_idle()
		}

		if !busy {
			time.Sleep(time.Duration(10) * time.Millisecond)
		}
	}

	self.on_stop()
}

func (self *Loop) do_update() {
	now := time.Now().Unix()
	if self.last != now {
		self.last = now

		// do something like as ACT schedule
	}
}

func (self *Loop) do_talk(talk *talk) {
	talk.s.DoPacket(talk.p)
}

func (self *Loop) do_idle() bool {
	return false
}

// ============================================================================

func (self *Loop) on_start() {
	schedule.Register(_LOOP_NAME, self)
}

func (self *Loop) on_stop() {
	schedule.UnRegister(_LOOP_NAME)
}

// ============================================================================

func (self *Loop) OnSchedule(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Loop) OnEvent(evt *event.Event) {
	event.Fire(evt.Id, evt.Args...)
}

func (self *Loop) OnTimer(id uint64) {
	fmt.Println("Loop.OnTimer:", id)
}

// ============================================================================
// public

func (self *Loop) PostEvent(evt *event.Event) {
	self.evtMgr.Fire(evt)
}

func (self *Loop) PostPacket(s iSession, p *tcp.Packet) {
	self.talks <- &talk{s, p}
}

func (self *Loop) CreateTimer(i uint64, r bool, f func()) uint64 {
	return self.timerMgr.CreateTimer(i, r, f)
}

func (self *Loop) CancelTimer(id uint64) {
	self.timerMgr.CancelTimer(id)
}

// ============================================================================

func Get() *Loop {
	return _loop
}

func Start() {
	if _loop == nil {
		return
	}

	_loop = &Loop{}
	_loop.start()
}

func Stop() {
	_loop.stop()
}
