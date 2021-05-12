package loop

import (
	"gworld/core/utils"
	"time"
)

var (
	_once   = []func(){}
	_update = []func(){}

	// please priority queue instead
	_timerq = []*timer_t{}

	_next      = []func(){}
	_next_once = map[string]func(){}
)

type timer_t struct {
	id     uint64
	fn     func()
	end    int64
	repeat bool
}

// ----------------------------------------------------------------------------
// export

func Run() {
	go func() {
		for {

			// once
			do_once()

			// timer
			do_timer()

			// update
			for _, fn := range _update {
				fn()
			}

			DoNext()

			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func Register(fn func()) {
	_update = append(_update, fn)
}

func Post(fn func()) {
	_once = append(_once, fn)
}

func SetTimer(fn func(), delay int64, repeat bool) {
	end := delay + time.Now().UnixNano()/1e6
	_timerq = append(_timerq, &timer_t{utils.SeqU64(), fn, end, repeat})
}

func ClearTimer(tid uint64) {
	for k, t := range _timerq {
		if t.id == tid {
			_timerq = append(_timerq[:k], _timerq[k+1:]...)
		}
	}
}

func Next(fn func()) {
	_next = append(_next, fn)
}

func NextOnce(key string, fn func()) {
	if _, ok := _next_once[key]; !ok {
		_next_once[key] = fn
	}
}

func DoNext() {
	if len(_next) > 0 {
		for _, fn := range _next {
			fn()
		}
		_next = _next[:0]
	}

	if len(_next_once) > 0 {
		for _, fn := range _next {
			fn()
		}
		_next_once = map[string]func(){}
	}
}

// ----------------------------------------------------------------------------
// local

func do_once() {
	if len(_once) == 0 {
		return
	}

	for _, fn := range _once {
		fn()
	}

	_once = _once[:0]
}

func do_timer() {
	if len(_timerq) == 0 {
		return
	}

	d := map[uint64]bool{}
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
