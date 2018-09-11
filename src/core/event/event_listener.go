package event

import (
	"sync/atomic"
)

var (
	_evts = map[string]map[int64]func(...interface{}){}
	_once = map[string][]func(...interface{}){}
)

var (
	_seq int64 = 1
)

// ---------------------------------------------------------------

func seq() int64 {
	return atomic.AddInt64(&_seq, 1)
}

// ---------------------------------------------------------------

func On(evt string, f func(...interface{})) int64 {
	if _evts[evt] == nil {
		_evts[evt] = make(map[int64]func(...interface{}))
	}

	seq := seq()

	_evts[evt][seq] = f

	return seq
}

func Cancel(evt string, seq int64) {
	if _evts[evt] != nil {
		delete(_evts[evt], seq)
	}
}

func Once(evt string, f func(...interface{})) {
	_once[evt] = append(_once[evt], f)
}

func Fire(evt string, args ...interface{}) {
	for _, f := range _evts[evt] {
		f(args...)
	}

	for _, f := range _once[evt] {
		f(args...)
	}

	_once[evt] = _once[evt][:0]
}
