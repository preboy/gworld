package thread

import (
	"sync"
	"time"
)

type Thread struct {
	u    func()
	run  bool
	quit chan bool
	w    *sync.WaitGroup
	itvl uint32
}

func (self *Thread) Go() {
	go func() {
		self.w.Add(1)
		defer func() {
			self.w.Done()
		}()

		if self.run {
			return
		}
		self.run = true
		defer func() {
			self.run = false
		}()

		for {
			select {
			case <-self.quit:
				return
			default:
				time.Sleep(time.Duration(self.itvl) * time.Millisecond)
			}
			self.u()
		}
	}()
}

func (self *Thread) Stop() {
	self.quit <- true
	self.w.Wait()
}

func (self *Thread) IsRunning() bool {
	return self.run
}

func NewThread(f func(), i uint32) *Thread {
	if f != nil {
		return &Thread{
			u:    f,
			w:    &sync.WaitGroup{},
			itvl: i,
			quit: make(chan bool),
		}
	}
	return nil
}
