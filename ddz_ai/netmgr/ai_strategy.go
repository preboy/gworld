package netmgr

import (
	"gworld/core/log"
	"gworld/ddz/lobby/poker"
	"math/rand"
)

type STRATEGY int

const (
	STRATEGY_CALC STRATEGY = 1 + iota // 分析
	STRATEGY_MUST                     // 能大就大
	STRATEGY_PASS                     // 能不出就不出

	STRATEGY_MAX
)

const (
	STRATEGY_METHOD = STRATEGY_MUST
)

var (
	_strategies [STRATEGY_MAX]strategy_t
)

type strategy_t struct {
	on_call func() int32
	on_play func(first bool) (cards []poker.Card)
}

func init() {

	// ------------------------------------------------------------------------
	// calc

	_strategies[STRATEGY_CALC].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_CALC].on_play = func(first bool) (cards []poker.Card) {
		a1 := poker.NewAnalyse(_ai.cards)
		a2 := poker.NewAnalyse(_ai.cards_left)

		// TODO  计算其它方位可能有的牌

		_ = a2

		if first {
			// 首出

		} else {

			// 跟牌

			cards_prev := _ai.prev_play()
			if cards_prev == nil {
				panic("cards_prev == nil")
			}

			// 上一首不为Nil的牌
			ci := poker.CardsAnalyse(cards_prev)

			// 找到大过他的牌
			ret := a1.Exceed(ci)
			for i, v := range ret {
				log.Info("the <%d> groups: %v", i, poker.CardsToString(v))
			}
		}

		return
	}

	// ------------------------------------------------------------------------
	// must

	_strategies[STRATEGY_MUST].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_MUST].on_play = func(first bool) (cards []poker.Card) {
		a1 := poker.NewAnalyse(_ai.cards)

		if first {
			// TODO 随机出一手，尽量能回收，切不把手里的牌搞乱
			c := cards_divide_abcde(_ai.cards)
			cards = c.first()
		} else {

			// 上一首不为Nil的牌
			cards_prev := _ai.prev_play()
			if cards_prev == nil {
				panic("cards_prev == nil")
			}

			ci := poker.CardsAnalyse(cards_prev)

			// 找到大过他的牌
			ret := a1.Exceed(ci)
			for i, v := range ret {
				log.Info("the <%d> Exceed er: %v", i, poker.CardsToString(v))
			}

			// TODO 处理带牌的情况

			if len(ret) > 0 {
				cards = ret[0]
			}
		}

		return
	}

	// ------------------------------------------------------------------------
	// pass

	_strategies[STRATEGY_PASS].on_call = func() int32 {
		return rand.Int31n(4)
	}

	_strategies[STRATEGY_PASS].on_play = func(first bool) (cards []poker.Card) {
		if first {
			poker.CardsSort(_ai.cards)
			l := len(_ai.cards)
			cards = _ai.cards[l-1:]
		}

		return
	}
}

// ----------------------------------------------------------------------------
// ai operation

func (ai *AILogic) ai_call() int32 {
	return _strategies[STRATEGY_METHOD].on_call()
}

func (ai *AILogic) ai_play(first bool) (cards []poker.Card) {
	return _strategies[STRATEGY_METHOD].on_play(first)
}
