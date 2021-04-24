package loop

import (
	"sync/atomic"
	"time"
)

var (
	_seq    = int64(1)
	_once   = []func(){}
	_update = []func(){}

	// please priority queue instead
	_timerq = []*timer_t{}
)

type timer_t struct {
	id     int64
	fn     func()
	end    int64
	repeat bool
}

// ----------------------------------------------------------------------------
// local

func new_seq() int64 {
	return atomic.AddInt64(&_seq, 1)
}

// ----------------------------------------------------------------------------
// export

func Run() {
	go func() {

		for {
			// post
			if len(_once) > 0 {
				for _, fn := range _once {
					fn()
				}
				_once = _once[:0]
			}

			// timer
			if len(_timerq) > 0 {
				d := map[int64]bool{}
				n := time.Now().UnixNano() / 1e6

				for _, t := range _timerq {
					if n >= t.end {
						t.fn()
						if !t.repeat {
							d[t.id] = true
						}
					}
				}

				// delete
				if len(d) > 0 {
					for k := range d {
						ClearTimer(k)
					}
				}
			}

			// update
			for _, fn := range _update {
				fn()
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func Post(fn func()) {
	_once = append(_once, fn)
}

func Register(fn func()) {
	_update = append(_update, fn)
}

func SetTimer(fn func(), delay int64, repeat bool) {
	end := delay + time.Now().UnixNano()/1e6
	_timerq = append(_timerq, &timer_t{new_seq(), fn, end, repeat})
}

func ClearTimer(tid int64) {
	for k, t := range _timerq {
		if t.id == tid {
			_timerq = append(_timerq[:k], _timerq[k+1:]...)
		}
	}
}
