package lobby

import (
	"gworld/core/log"
)

type stage_func struct {
	OnEnter  func(m *Match)
	OnUpdate func(m *Match)
	OnLeave  func(m *Match)
}

var (
	FSM [stage_MAX]stage_func
)

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
		m.InitDeck()

		// 发牌
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
		m.first_call = next_seat()
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
