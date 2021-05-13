package smatch

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby/poker"
	"gworld/ddz/loop"
	"gworld/ddz/pb"
	"time"
)

type stage_func struct {
	OnEnter   func(t *Table)
	OnLeave   func(t *Table)
	OnUpdate  func(t *Table)
	OnMessage func(t *Table, pid string, req comp.IMessage, res comp.IMessage)
}

var (
	FSM [stage_max]stage_func
)

type deck_info_t struct {
	index int32
	start int64

	deal_info []*deal_info_t
	call_info []*call_info_t
	caca_info *cacl_info_t
	play_info []*play_info_t
	calc_info *calc_info_t
}

type deal_info_t struct {
	pos   seat
	cards []poker.Card
}

type call_info_t struct {
	past  int64
	pos   seat
	score int32 // 0,1,2,3
}

type cacl_info_t struct {
	draw  bool
	lord  seat
	score int32
}

type play_info_t struct {
	past  int
	pos   seat
	cards []poker.Card // empty is PASS
}

// 结算信息
type calc_info_t struct {
	win    seat // -1 流局
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

	FSM[stage_wait].OnEnter = func(t *Table) {
		log.Info("enter prepare")
	}

	FSM[stage_wait].OnLeave = func(t *Table) {
		log.Info("leave prepare")
	}

	FSM[stage_wait].OnUpdate = func(t *Table) {
		// 坐满之后自动开启
		full := true
		for i := seat_east; i < seat_max; i++ {
			if t.seats[i] == nil {
				full = false
				break
			}
		}

		if full {
			t.Switch(stage_deal)
		}
	}

	FSM[stage_wait].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {
	}

	// ------------------------------------------------------------------------
	// deal

	FSM[stage_deal].OnEnter = func(t *Table) {
		log.Info("enter deal")
		t.DeckOpen()

		// 发牌
		{
			t.cards = poker.CreatePoker()

			n := t.call_pos
			t.DealCards(n, t.cards[:17])

			n = next_seat(n)
			t.DealCards(n, t.cards[17:34])

			n = next_seat(n)
			t.DealCards(n, t.cards[34:51])
		}
	}

	FSM[stage_deal].OnLeave = func(t *Table) {
		log.Info("leave deal")
	}

	FSM[stage_deal].OnUpdate = func(t *Table) {
		t.Switch(stage_call)
	}

	FSM[stage_deal].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {
	}

	// ------------------------------------------------------------------------
	// call

	FSM[stage_call].OnEnter = func(t *Table) {
		log.Info("enter call")

		t.SendActionCall(t.call_pos)
	}

	FSM[stage_call].OnLeave = func(t *Table) {
		// 发送叫分结果
		msg := &pb.CallScoreCalcBroadcast{
			Draw:  t.deck_info.caca_info.draw,
			Lord:  int32(t.deck_info.caca_info.lord),
			Score: t.deck_info.caca_info.score,
		}

		t.seats[t.deck_info.caca_info.lord].AddCards(t.cards[51:])

		for _, v := range t.cards[51:] {
			msg.Cards = append(msg.Cards, v.Value())
		}

		t.Broadcast(msg)
		log.Info("leave call")
	}

	FSM[stage_call].OnUpdate = func(t *Table) {
		// 时间到了，叫0分
		if time.Since(t.action_ts) > 15*time.Second {
			t.deck_info.call_info = append(t.deck_info.call_info, &call_info_t{15, t.call_pos, 0})
			t.SendActionCall(next_seat(t.call_pos))
		}
	}

	FSM[stage_call].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {

		switch req.GetOP() {
		case pb.Default_CallScoreRequest_OP:
			{
				r := req.(*pb.CallScoreRequest)
				s := res.(*pb.CallScoreResponse)

				// 收到消息
				pos, ok := t.pid_to_pos(pid)
				if !ok {
					s.ErrCode = gconst.Err_CallPid
					return
				}

				if pos != t.call_pos {
					s.ErrCode = gconst.Err_CallPos
					return
				}

				if r.Score < 0 && r.Score > 3 {
					s.ErrCode = gconst.Err_CallScore
					return
				}

				// 检测分数是否合法
				for _, v := range t.deck_info.call_info {
					if v.score > 0 && r.Score <= v.score {
						// s.ErrCode = gconst.Err_CallScore2
						// return
						r.Score = 0
					}
				}

				delay := time.Since(t.action_ts).Seconds()
				t.deck_info.call_info = append(t.deck_info.call_info, &call_info_t{int64(delay), pos, r.Score})

				s.ErrCode = gconst.Err_OK

				loop.Next(func() {
					t.Broadcast(&pb.CallScoreResultBroadcast{
						Pos:   int32(pos),
						Score: r.Score,
					})

					t.SendActionCall(next_seat(t.call_pos))
				})
			}
		default:
			break
		}
	}

	// ------------------------------------------------------------------------
	// play

	FSM[stage_play].OnEnter = func(t *Table) {
		log.Info("enter play")

		t.play_pos = t.deck_info.caca_info.lord
		t.play_idx = 0
		t.play_pass = 0

		t.action_ts = time.Now()

		t.Broadcast(&pb.PlayBroadcast{
			Pos:   int32(t.play_pos),
			First: t.play_idx == 0,
		})
	}

	FSM[stage_play].OnLeave = func(t *Table) {
		log.Info("leave play")
	}

	FSM[stage_play].OnUpdate = func(t *Table) {
		// 默认出牌
		if time.Since(t.action_ts) < 15*time.Second {
			return
		}

		req := &pb.PlayRequest{}
		res := &pb.PlayResponse{}

		pid, ok := t.pos_to_pid(t.play_pos)
		if !ok {
			panic("Invalid play pos")
		}

		// 首家出最小牌 非首家不出牌
		if t.play_idx == 0 {
			cards := t.seats[t.play_pos].GetDefaultCards()
			req.Cards = poker.CardsToInt32(cards)
		}

		FSM[stage_play].OnMessage(t, pid, req, res)
	}

	FSM[stage_play].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {

		switch req.GetOP() {
		case pb.Default_PlayRequest_OP:
			{
				r := req.(*pb.PlayRequest)
				s := res.(*pb.PlayResponse)

				pos, ok := t.pid_to_pos(pid)
				if !ok {
					panic("invalid pid to play")
				}

				if pos != t.play_pos {
					s.ErrCode = gconst.Err_NotYourTurn
					return
				}

				var valid bool
				var cards []poker.Card
				var ci *poker.CardsInfo = &poker.CardsInfo{}

				if len(r.Cards) != 0 {
					// 是否合法的牌
					cards, valid = poker.CardsFromInt32(r.Cards)
					if !valid {
						s.ErrCode = gconst.Err_CardInvalid
						return
					}

					// 是否手上有这些牌
					if !t.seats[pos].ExistCards(cards) {
						s.ErrCode = gconst.Err_CardNotExist
						return
					}

					ci = poker.CardsAnalyse(cards)
				}

				// 牌型检测
				if t.play_idx == 0 {
					// 首家不能为空
					if len(r.Cards) == 0 {
						s.ErrCode = gconst.Err_CardNull
						return
					}

					// 是否合法的牌型
					if ci.Type == poker.CardsTypeNIL {
						s.ErrCode = gconst.Err_CardTypeInvalid
						return
					}

					t.PlayHand(cards, ci)
				} else {
					if len(r.Cards) == 0 {
						t.PlayPass()
					} else {
						// 是否合法的牌型
						if ci.Type != t.play_cards.Type || ci.Max <= t.play_cards.Max || ci.Len != t.play_cards.Len {
							s.ErrCode = gconst.Err_CardTypeInvalid
							return
						}

						t.PlayHand(cards, ci)
					}
				}

				s.ErrCode = gconst.Err_OK
			}
		default:
			break
		}

	}

	// ------------------------------------------------------------------------
	// calc

	FSM[stage_calc].OnEnter = func(t *Table) {
		log.Info("enter calc")

		msg := &pb.DeckEndBroadcast{}

		// 结算信息

		lord_win := false
		if t.play_pos == t.deck_info.caca_info.lord {
			lord_win = true
		}

		base_score := t.deck_info.caca_info.score
		if t.bombs > 0 {
			base_score *= t.bombs
		}

		for k := range t.seats {
			pos := int32(k)
			score := base_score

			if pos == t.deck_info.caca_info.lord {
				score *= 2

				if !lord_win {
					score = 0 - score
				}
			} else {

				if lord_win {
					score = 0 - score
				}
			}

			msg.Score = append(msg.Score, score)
		}

		t.Broadcast(msg)
	}

	FSM[stage_calc].OnLeave = func(t *Table) {
		t.DeckClosed()
		log.Info("leave calc")
	}

	FSM[stage_calc].OnUpdate = func(t *Table) {
		t.NextDeck()
	}

	FSM[stage_calc].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {
	}

	// ------------------------------------------------------------------------
	// over

	FSM[stage_over].OnEnter = func(t *Table) {
		log.Info("enter over")
	}

	FSM[stage_over].OnLeave = func(t *Table) {
		log.Info("leave over")
	}

	FSM[stage_over].OnUpdate = func(t *Table) {

	}

	FSM[stage_over].OnMessage = func(t *Table, pid string, req comp.IMessage, res comp.IMessage) {
	}
}

// ============================================================================
// local

func next_seat(seat seat) seat {
	seat++

	if seat >= seat_max {
		seat = seat_east
	}

	return seat
}
