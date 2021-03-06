package timer

import (
	"gworld/core/thread"
	"gworld/core/utils"

	"sync"
	"time"
)

var (
	_curr_millisec uint64
	_lock          *sync.RWMutex
	_thread        *thread.Thread
	_now           time.Time
)

type Timer struct {
	id       uint64
	fn       func()
	curr     uint64
	repeat   bool
	interval uint64
}

type TimerMgr struct {
	timers   map[uint64]*Timer
	receiver ITimerMgr
}

type ITimerMgr interface {
	OnTimer(id uint64)
}

func NewTimerMgr(r ITimerMgr) *TimerMgr {
	return &TimerMgr{
		timers:   make(map[uint64]*Timer),
		receiver: r,
	}
}

// may be run in other goroutine
func (self *TimerMgr) Update() (busy bool) {
	now := get_tick()
	for id, timer := range self.timers {
		if now >= timer.curr+timer.interval {
			if timer.fn != nil {
				timer.fn()
			} else {
				self.receiver.OnTimer(id)
			}
			if timer.repeat {
				timer.curr = now
			} else {
				self.CancelTimer(id)
			}
			busy = true
		}
	}
	return
}

func (self *TimerMgr) CreateTimer(i uint64, r bool, f func()) uint64 {
	id := utils.SeqU64()
	timer := &Timer{
		id:       id,
		interval: i,
		repeat:   r,
		fn:       f,
		curr:     get_tick(),
	}
	self.timers[id] = timer
	return id
}

func (self *TimerMgr) CancelTimer(id uint64) {
	// It's safe to remove key for map even if in range
	delete(self.timers, id)
}

// ============================================================================
// package method

func get_tick() uint64 {
	_lock.RLock()
	defer func() {
		_lock.RUnlock()
	}()
	return _curr_millisec
}

func update_tick() {
	_lock.Lock()
	_curr_millisec = uint64(float64(time.Since(_now).Nanoseconds()) / 1000 / 1000)
	_lock.Unlock()
}

func Start() {
	_lock = &sync.RWMutex{}
	if _thread == nil {
		_now = time.Now()
		update_tick()
		_thread = thread.NewThread(update_tick, 20)
		_thread.Go()
	}
}

func Stop() {
	if _thread != nil {
		_thread.Stop()
	}
}
