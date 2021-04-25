package utils

import "sync/atomic"

var (
	_seq32  int32  = 0
	_seq64  int64  = 0
	_sequ32 uint32 = 0
	_sequ64 uint64 = 0
)

func Seq32() int32 {
	return atomic.AddInt32(&_seq32, 1)
}

func Seq64() int64 {
	return atomic.AddInt64(&_seq64, 1)
}

func SeqU32() uint32 {
	return atomic.AddUint32(&_sequ32, 1)
}

func SeqU64() uint64 {
	return atomic.AddUint64(&_sequ64, 1)
}
