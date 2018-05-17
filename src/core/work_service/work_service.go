package work_service

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	__WorkService *WorkService
)

type WorkService struct {
	n       int32
	w       *sync.WaitGroup
	l       *sync.Mutex
	c       *sync.Cond
	q       chan func()
	running bool
}

func NewWorkService() *WorkService {
	s := &WorkService{
		w: &sync.WaitGroup{},
		l: &sync.Mutex{},
		q: make(chan func(), 2048),
	}
	s.c = sync.NewCond(s.l)
	return s
}

func (self *WorkService) Start(n int) bool {
	if self.running || n <= 0 || n > 64 {
		return false
	}

	self.running = true
	for i := 0; i < n; i++ {
		i := i
		atomic.AddInt32(&self.n, 1)

		go func(i int) {
			defer func() {
				self.w.Done()
				atomic.AddInt32(&self.n, -1)
			}()

			self.w.Add(1)

			for {
				if !self.running {
					return
				}

				self.c.L.Lock()
				self.c.Wait()

				// pick one function from q     // all goroutine exclude
				select {
				case f, ok := <-self.q:
					self.c.L.Unlock()
					if !ok {
						return
					} else {
						f() // feel happiness self
					}
				default:
					self.c.L.Unlock()
				}

			}

		}(i)
	}

	// monitor
	go func() {
		defer self.w.Done()
		self.w.Add(1)

		for {
			if atomic.LoadInt32(&self.n) == 0 {
				return
			}
			if len(self.q) > 0 {
				self.c.Signal()
				time.Sleep(1000 * time.Microsecond)
			}
		}
	}()

	return true
}

func (self *WorkService) Stop() {
	if !self.running {
		return
	}

	self.running = false
	close(self.q)

	self.c.Broadcast()
	self.w.Wait()

	fmt.Println("Closeing..........")

	// No orphan func
	for f := range self.q {
		f()
	}
}

func (self *WorkService) Queue(f func()) {
	if self.running {
		self.q <- f
		self.c.Signal()
	}
}

func Start(n int) {
	__WorkService = NewWorkService()
	__WorkService.Start(n)
}

func Stop() {
	__WorkService.Stop()
}

func Queue(f func()) {
	__WorkService.Queue(f)
}

// func main() {

// 	s := NewWorkService()
// 	s.Start(10)

// 	for i := 0; i < 70; i++ {
// 		i := i
// 		s.Queue(func() {
// 			fmt.Println(i, "start")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println(i, "end")
// 		})
// 	}

// 	time.Sleep(6 * time.Second)
// 	fmt.Println("read to end ----------------")

// 	s.Close()
// }
