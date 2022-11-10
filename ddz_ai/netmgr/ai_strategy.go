package netmgr

import "math/rand"

type STRATEGY int

const (
	STRATEGY_CALC STRATEGY = 1 + iota // 分析
	STRATEGY_MUST                     // 能大就大
	STRATEGY_PASS                     // 能不出就不出

	STRATEGY_MAX
)

type strategy_t struct {
	on_call func() int32
	on_play func()
}

var (
	_strategies [STRATEGY_MAX]strategy_t
)

func init() {

	// calc
	_strategies[STRATEGY_CALC].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_CALC].on_play = func() {

	}

	// must
	_strategies[STRATEGY_MUST].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_MUST].on_play = func() {

	}

	// pass
	_strategies[STRATEGY_PASS].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_PASS].on_play = func() {

	}
}
