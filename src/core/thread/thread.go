package thread

import (
	"time"
)

type thread_func = func()

type Thread struct {
	u    thread_func
	run  bool
	quit chan int
}

func (self *Thread) Go() {
	go func() {
		if self.run {
			return
		}
		self.run = true
		defer func() {
			self.run = false
		}()
		for {
			select {
			case <-quit:
				break
			default:
				time.Sleep(20 * time.Millisecond)
			}
			self.u()
		}
	}()
}

func (self *Thread) Stop() {
	self.quit <- 0
}

func (self *Thread) IsRunning() {
	return self.run
}

func NewThread(f thread_func) *Thread {
	if f != nil {
		return &Thread{
			u: f,
		}
	}
	return nil
}
