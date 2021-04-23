package loop

import (
	"time"
)

var (
	_once  = []func(){}
	_funcs = []func(){}
)

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

			// update
			for _, fn := range _funcs {
				fn()
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func Register(fn func()) {
	_funcs = append(_funcs, fn)
}

func Post(fn func()) {
	_once = append(_once, fn)
}
