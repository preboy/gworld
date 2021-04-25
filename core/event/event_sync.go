package event

import (
	"gworld/core/utils"
)

var (
	_evts = map[uint32]map[uint64]func(uint32, ...interface{}){}
	_once = map[uint32][]func(uint32, ...interface{}){}
)

// ============================================================================

func On(id uint32, fn func(uint32, ...interface{})) uint64 {
	if _evts[id] == nil {
		_evts[id] = make(map[uint64]func(uint32, ...interface{}))
	}

	seq := utils.SeqU64()
	_evts[id][seq] = fn

	return seq
}

func Cancel(id uint32, seq uint64) {
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
