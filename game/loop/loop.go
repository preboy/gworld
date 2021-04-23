package loop

// the main loop in game

import (
	"sync"
	"time"

	"gworld/core/event"
	"gworld/core/log"
	"gworld/core/schedule"
	"gworld/core/tcp"
	"gworld/core/timer"
	"gworld/core/utils"
	"gworld/game/modules/act"
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
	funcs    []func()
	idles    []func()
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
	close(self.q)
	self.w.Wait()
}

// ============================================================================

func (self *Loop) working() {
	self.w.Add(1)

	defer func() {
		self.w.Done()
	}()

	self.on_start()

	go func() {

		func() {
			var busy bool

			for {
				busy = false

				select {
				case <-self.q:
					return

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

				if self.do_update_secondly() {
					busy = true
				}

				if self.do_funcs() {
					busy = true
				}

				if !busy {
					busy = self.do_idle()
				}

				if !busy {
					time.Sleep(time.Duration(10) * time.Millisecond)
				}
			}
		}()

		self.on_stop()
	}()
}

func (self *Loop) do_update_secondly() bool {
	now := time.Now().Unix()

	if self.last != now {
		self.last = now

		utils.ExecuteSafely(func() {
			// do something as ACT schedule
			act.OnLoopUpdate()
		})

		return true
	}

	return false
}

func (self *Loop) do_idle() bool {
	busy := false

	if len(self.idles) > 0 {
		for _, fn := range self.idles {
			utils.ExecuteSafely(func() {
				fn()
				busy = true
			})
		}
		self.idles = self.idles[:0]
	}

	return busy
}

func (self *Loop) do_talk(talk *talk) {
	if talk != nil {
		talk.s.DoPacket(talk.p)
	}
}

func (self *Loop) do_funcs() bool {
	if len(self.funcs) == 0 {
		return false
	}

	for _, fn := range self.funcs {
		utils.ExecuteSafely(func() {
			fn()
		})
	}

	self.funcs = self.funcs[:0]

	return true
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
	log.Info("Loop.OnTimer:", id)
}

// ============================================================================
// public

func (self *Loop) PostEventArgs(id uint32, args ...interface{}) {
	evt := event.NewEvent(id, args...)
	self.PostEvent(evt)
}

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

func (self *Loop) PostFunc(fn func()) {
	self.funcs = append(self.funcs, fn)
}

func (self *Loop) PostIdle(fn func()) {
	self.idles = append(self.idles, fn)
}

// ============================================================================

func Get() *Loop {
	return _loop
}

func Start() {
	if _loop != nil {
		return
	}

	_loop = &Loop{}
	_loop.start()
}

func Stop() {
	_loop.stop()
}
