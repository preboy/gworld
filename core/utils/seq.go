package utils

import "sync/atomic"

var (
	_sequ32 uint32 = 1
	_sequ64 uint64 = 1
)

func SeqU32() uint32 {
	return atomic.AddUint32(&_sequ32, 1)
}

func SeqU64() uint64 {
	return atomic.AddUint64(&_sequ64, 1)
}
