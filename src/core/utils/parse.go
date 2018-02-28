package utils

import (
	"strconv"
)

func I32toa(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func U32toa(n uint32) string {
	return strconv.FormatUint(uint64(n), 10)
}

func I64toa(n int64) string {
	return strconv.FormatInt(n, 10)
}

func U64toa(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func Atoi32(v string) int32 {
	n, err := strconv.ParseInt(v, 10, 32)
	if err == nil {
		return int32(n)
	} else {
		return 0
	}
}

func Atou32(v string) uint32 {
	n, err := strconv.ParseUint(v, 10, 32)
	if err == nil {
		return uint32(n)
	} else {
		return 0
	}
}

func Atoi64(v string) int64 {
	n, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		return n
	} else {
		return 0
	}
}

func Atou64(v string) uint64 {
	n, err := strconv.ParseUint(v, 10, 64)
	if err == nil {
		return n
	} else {
		return 0
	}
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
