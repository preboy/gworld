package event

import (
	"core/utils"
	"sync/atomic"
)

var (
	_evts = map[uint32]map[int64]func(uint32, ...interface{}){}
	_once = map[uint32][]func(uint32, ...interface{}){}
)

var (
	_seq int64 = 1
)

// ============================================================================

func seq() int64 {
	return atomic.AddInt64(&_seq, 1)
}

// ============================================================================

func On(id uint32, fn func(uint32, ...interface{})) int64 {
	if _evts[id] == nil {
		_evts[id] = make(map[int64]func(uint32, ...interface{}))
	}

	seq := seq()

	_evts[id][seq] = fn

	return seq
}

func Cancel(id uint32, seq int64) {
	if _evts[id] != nil {
		delete(_evts[id], seq)
	}
}

func Once(id uint32, fn func(uint32, ...interface{})) {
	_once[id] = append(_once[id], fn)
}

func Fire(id uint32, args ...interface{}) {
	for _, fn := range _evts[id] {
		utils.ExecuteSafely(func() {
			fn(id, args...)
		})
	}

	for _, fn := range _once[id] {
		utils.ExecuteSafely(func() {
			fn(id, args...)
		})
	}

	_once[id] = _once[id][:0]
}
