package lobby

import (
	"gworld/core/log"
)

type stage_func struct {
	OnEnter   func(m *Match)
	OnUpdate  func(m *Match)
	OnMessage func(m *Match)
	OnLeave   func(m *Match)
}

var (
	FSM [stage_MAX]stage_func
)

type deck_info_t struct {
	index int
	start int64

	deal  []*deal_info_t
	call  []*call_info_t
	callr *call_result_t
	play  []*play_info_t
	calc  *calc_info_t
}

type deal_info_t struct {
	pos   SEAT
	cards []Card
}

type call_info_t struct {
	past  int64
	pos   SEAT
	score int // 0,1,2,3
}

type call_result_t struct {
	lord  SEAT
	score int
}

type play_info_t struct {
	past  int
	pos   SEAT
	cards []Card // empty is PASS
}

// 结算信息
type calc_info_t struct {
	win    SEAT // -1 流局
	lord   bool
	score  int
	spring bool
	bomb   int
}

// ----------------------------------------------------------------------------
// init

func init() {
	// ------------------------------------------------------------------------
	// prepare

	FSM[stage_prepare].OnEnter = func(m *Match) {
		log.Info("enter prepare")
	}

	FSM[stage_prepare].OnUpdate = func(m *Match) {
		m.Switch(stage_deal)
	}

	FSM[stage_prepare].OnLeave = func(m *Match) {
		log.Info("leave prepare")
	}

	// ------------------------------------------------------------------------
	// deal

	FSM[stage_deal].OnEnter = func(m *Match) {
		log.Info("enter deal")
		m.DeckOpen()

		// 发牌
		{
			cards := NewPoker()

			n := m.call_pos
			m.seats[n].AddCards(cards[:17])
			m.deck_data.deal = append(m.deck_data.deal, &deal_info_t{n, cards[:17]})

			n = next_seat(n)
			m.seats[n].AddCards(cards[17:34])
			m.deck_data.deal = append(m.deck_data.deal, &deal_info_t{n, cards[17:34]})

			n = next_seat(n)
			m.seats[n].AddCards(cards[34:51])
			m.deck_data.deal = append(m.deck_data.deal, &deal_info_t{n, cards[34:51]})
		}
	}

	FSM[stage_deal].OnUpdate = func(m *Match) {
		m.Switch(stage_call)
	}

	FSM[stage_deal].OnLeave = func(m *Match) {
		log.Info("leave deal")
	}

	// ------------------------------------------------------------------------
	// call

	FSM[stage_call].OnEnter = func(m *Match) {
		log.Info("enter call")

		m.SetActionCall(m.call_pos)
	}

	FSM[stage_call].OnUpdate = func(m *Match) {

	}

	FSM[stage_call].OnLeave = func(m *Match) {
		log.Info("leave call")
	}

	// ------------------------------------------------------------------------
	// play

	FSM[stage_play].OnEnter = func(m *Match) {
		log.Info("enter play")
	}

	FSM[stage_play].OnUpdate = func(m *Match) {

	}

	FSM[stage_play].OnLeave = func(m *Match) {
		log.Info("leave play")
	}

	// ------------------------------------------------------------------------
	// calc

	FSM[stage_calc].OnEnter = func(m *Match) {
		log.Info("enter calc")
	}

	FSM[stage_calc].OnUpdate = func(m *Match) {
		m.NextDeck()
	}

	FSM[stage_calc].OnLeave = func(m *Match) {
		m.DeckClosed()
		log.Info("leave calc")
	}

	// ------------------------------------------------------------------------
	// over

	FSM[stage_over].OnEnter = func(m *Match) {
		log.Info("enter over")
	}

	FSM[stage_over].OnUpdate = func(m *Match) {

	}

	FSM[stage_over].OnLeave = func(m *Match) {
		log.Info("leave over")
	}
}

// ============================================================================
// local

func next_seat(seat SEAT) SEAT {
	seat++

	if seat == SEAT_MAX {
		seat = SEAT_EAST
	}

	return seat
}
