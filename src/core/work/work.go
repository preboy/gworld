package work

import (
	"sync"
	"sync/atomic"
	"time"

	"game/loop"
)

var (
	mgr *work_mgr
)

type work struct {
	fn func() func()
	cb func()
}

type work_mgr struct {
	n     int32
	w     *sync.WaitGroup
	l     *sync.Mutex
	c     *sync.Cond
	works chan *work
}

func (self *work) do() {
	cb := self.fn()

	if cb != nil {
		loop.Get().PostCallback(cb)
	}

	if self.cb != nil {
		loop.Get().PostCallback(self.cb)
	}
}

func (self *work_mgr) start(n int) bool {
	if n <= 0 || n > 64 {
		return false
	}

	for i := 0; i < n; i++ {
		i := i

		atomic.AddInt32(&self.n, 1)

		go func(j int) {

			self.w.Add(1)

			defer func() {
				self.w.Done()
				atomic.AddInt32(&self.n, -1)
			}()

			for {

				self.c.L.Lock()
				self.c.Wait()

				select {
				case work, ok := <-self.works:
					self.c.L.Unlock()
					if !ok {
						return
					} else {
						work.do()
					}
				default:
					self.c.L.Unlock()
				}
			}
		}(i)
	}

	// monitor
	go func() {
		self.w.Add(1)
		defer self.w.Done()

		for {
			if atomic.LoadInt32(&self.n) == 0 {
				return
			}

			for len(self.works) > 0 {
				self.c.Signal()
			}

			time.Sleep(10 * 1000 * time.Microsecond)
		}
	}()

	return true
}

func (self *work_mgr) stop() {
	close(self.works)

	self.c.Broadcast()
	self.w.Wait()
}

func (self *work_mgr) queue(w *work) {
	self.works <- w
	self.c.Signal()
}

// ============================================================================

func Start(n int) {
	if mgr != nil {
		return
	}

	mgr = &work_mgr{
		w:     &sync.WaitGroup{},
		l:     &sync.Mutex{},
		works: make(chan *work, 0x1000),
	}

	mgr.c = sync.NewCond(mgr.l)
	mgr.start(n)
}

func Stop() {
	if mgr != nil {
		mgr.stop()
		mgr = nil
	}
}

func Queue(fn func() func(), cb func()) {
	if mgr != nil && fn != nil {
		mgr.queue(&work{fn, cb})
	}
}
