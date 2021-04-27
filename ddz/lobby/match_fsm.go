package lobby

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/pb"
	"time"
)

type stage_func struct {
	OnEnter   func(m *Match)
	OnLeave   func(m *Match)
	OnUpdate  func(m *Match)
	OnMessage func(m *Match, pid string, req comp.Message, res comp.Message)
}

var (
	FSM [stage_MAX]stage_func
)

type deck_info_t struct {
	index int
	start int64

	deal_info []*deal_info_t
	call_info []*call_info_t
	caca_info *call_calc_info_t
	play_info []*play_info_t
	calc_info *calc_info_t
}

type deal_info_t struct {
	pos   SEAT
	cards []Card
}

type call_info_t struct {
	past  int64
	pos   SEAT
	score int32 // 0,1,2,3
}

type call_calc_info_t struct {
	draw  bool
	lord  SEAT
	score int32
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
	score  int32
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

	FSM[stage_prepare].OnLeave = func(m *Match) {
		log.Info("leave prepare")
	}

	FSM[stage_prepare].OnUpdate = func(m *Match) {
		m.Switch(stage_deal)
	}

	FSM[stage_prepare].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {
	}

	// ------------------------------------------------------------------------
	// deal

	FSM[stage_deal].OnEnter = func(m *Match) {
		log.Info("enter deal")
		m.DeckOpen()

		// 发牌
		{
			m.cards = NewPoker()

			n := m.call_pos
			m.DealCards(n, m.cards[:17])

			n = next_seat(n)
			m.DealCards(n, m.cards[17:34])

			n = next_seat(n)
			m.DealCards(n, m.cards[34:51])
		}
	}

	FSM[stage_deal].OnLeave = func(m *Match) {
		log.Info("leave deal")
	}

	FSM[stage_deal].OnUpdate = func(m *Match) {
		m.Switch(stage_call)
	}

	FSM[stage_deal].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {
	}

	// ------------------------------------------------------------------------
	// call

	FSM[stage_call].OnEnter = func(m *Match) {
		log.Info("enter call")

		m.SendActionCall(m.call_pos)
	}

	FSM[stage_call].OnLeave = func(m *Match) {
		// 发送叫分结果
		msg := &pb.CallScoreCalcBroadcast{
			Draw:  m.deck_info.caca_info.draw,
			Lord:  int32(m.deck_info.caca_info.lord),
			Score: m.deck_info.caca_info.score,
		}

		for _, v := range m.cards[51:] {
			msg.Cards = append(msg.Cards, int32(v))
		}

		m.Broadcast(msg)
		log.Info("leave call")
	}

	FSM[stage_call].OnUpdate = func(m *Match) {
		// 时间到了，叫0分
		if time.Since(m.action_ts) > 15*time.Second {
			m.deck_info.call_info = append(m.deck_info.call_info, &call_info_t{15, m.call_pos, 0})
			m.SendActionCall(next_seat(m.call_pos))
		}
	}

	FSM[stage_call].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {

	}

	// ------------------------------------------------------------------------
	// play

	FSM[stage_play].OnEnter = func(m *Match) {
		log.Info("enter play")
	}

	FSM[stage_play].OnLeave = func(m *Match) {
		log.Info("leave play")
	}

	FSM[stage_play].OnUpdate = func(m *Match) {

	}

	FSM[stage_play].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {
	}

	// ------------------------------------------------------------------------
	// calc

	FSM[stage_calc].OnEnter = func(m *Match) {
		log.Info("enter calc")
	}

	FSM[stage_calc].OnLeave = func(m *Match) {
		m.DeckClosed()
		log.Info("leave calc")
	}

	FSM[stage_calc].OnUpdate = func(m *Match) {
		m.NextDeck()
	}

	FSM[stage_calc].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {
	}

	// ------------------------------------------------------------------------
	// over

	FSM[stage_over].OnEnter = func(m *Match) {
		log.Info("enter over")
	}

	FSM[stage_over].OnLeave = func(m *Match) {
		log.Info("leave over")
	}

	FSM[stage_over].OnUpdate = func(m *Match) {

	}

	FSM[stage_over].OnMessage = func(m *Match, pid string, req comp.Message, res comp.Message) {
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
