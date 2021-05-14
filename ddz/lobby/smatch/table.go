package smatch

import (
	"fmt"
	"time"

	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/lobby/poker"
	"gworld/ddz/loop"
	"gworld/ddz/pb"
)

type (
	seat  = int32
	stage = int32
)

const (
	seat_east seat = iota + 0
	seat_south
	seat_west
	seat_max
)

const (
	stage_wait stage = iota + 0 // 未开始
	stage_deal                  // 发牌
	stage_call                  // 叫分
	stage_play                  // 出牌
	stage_calc                  // 结算
	stage_over                  // 已结束
	stage_max
)

type Table struct {
	m *SMatch
	l *log.Logger

	id int32

	seats [seat_max]*gambler_table_t // 3个方位的pid

	stage stage

	deck_index int32 // 当前牌副数
	host_pos   seat  // 首叫方位

	// call
	call_pos seat // 叫分方位

	// play
	play_pos   seat             // 当前出牌的位置
	play_idx   int32            // 出牌顺序(round)
	play_pass  int32            // pass数量
	play_cards *poker.CardsInfo // 牌型

	bombs int32        // 炸弹数
	cards []poker.Card // 本副牌

	action_ts time.Time

	deck_info         *deck_info_t   // 当前副数据
	deck_info_history []*deck_info_t // 历史数据
}

func (self *Table) OnUpdate() {
	FSM[self.stage].OnUpdate(self)
	loop.DoNext()
}

func (self *Table) OnMessage(pid string, req comp.IMessage, res comp.IMessage) {
	FSM[self.stage].OnMessage(self, pid, req, res)
}

func (self *Table) Init(id int32, m *SMatch) {
	self.m = m
	self.id = id
	self.stage = stage_wait
	self.host_pos = seat_east
	self.deck_index = 0

	logname := fmt.Sprintf("ddz_%v_(%v)_%v.log", m.GetMID(), m.GetName(), self.id)
	self.l = log.NewLogger(logname, false)
}

func (self *Table) Release() {
	self.l.Stop()
}

func (self *Table) Sit(pid string) bool {
	join := false

	for i := seat_east; i < seat_max; i++ {
		if self.seats[i] != nil {
			continue
		}

		self.seats[i] = &gambler_table_t{
			m:    self,
			pid:  pid,
			pos:  i,
			data: &deck_data{},
		}

		join = true
		break
	}

	return join
}

func (self *Table) Switch(stage stage) {
	FSM[self.stage].OnLeave(self)
	self.stage = stage
	FSM[self.stage].OnEnter(self)
}

func (self *Table) IsOver() bool {
	return self.stage == stage_over
}

func (self *Table) DeckOpen() {
	self.deck_index++
	self.bombs = 0

	for i := seat_east; i < seat_max; i++ {
		self.seats[i].data = &deck_data{}
	}

	self.deck_info = &deck_info_t{
		index:     self.deck_index,
		start:     time.Now().Unix(),
		deal_info: nil,
		call_info: nil,
		caca_info: &cacl_info_t{},
		play_info: nil,
		calc_info: &calc_info_t{},
	}

	self.call_pos = self.host_pos
}

func (self *Table) DeckClosed() {
	self.host_pos = next_seat(self.host_pos)
	self.deck_info_history = append(self.deck_info_history, self.deck_info)
}

func (self *Table) NextDeck() {
	if self.deck_index < self.m.conf.TotalDeck {
		self.Switch(stage_deal)
	} else {
		self.Switch(stage_over)
	}
}

func (self *Table) Broadcast(msg comp.IMessage) {
	for _, v := range self.seats {
		v.SendMessage(msg)
	}
}

func (self *Table) Notify(pid string, msg comp.IMessage) {
	for _, v := range self.seats {
		if v.pid == pid {
			v.SendMessage(msg)
			break
		}
	}
}

func (self *Table) pos_to_pid(pos seat) (string, bool) {
	for _, v := range self.seats {
		if v.pos == pos {
			return v.pid, true
		}
	}

	return "", false
}

func (self *Table) pid_to_pos(pid string) (seat, bool) {
	for _, v := range self.seats {
		if v.pid == pid {
			return v.pos, true
		}
	}

	return seat_max, false
}

func (self *Table) DealCards(pos seat, cards []poker.Card) {

	self.seats[pos].AddCards(cards)
	self.l.Info("%v 拿到的手牌为：%v", pos_to_string(pos), poker.CardsToString(self.seats[pos].data.cards))

	self.deck_info.deal_info = append(self.deck_info.deal_info, &deal_info_t{pos, cards})

	msg := &pb.DealCardNotify{
		Pos: int32(pos),
	}

	for _, v := range cards {
		msg.Cards = append(msg.Cards, v.Value())
	}

	if pid, ok := self.pos_to_pid(pos); ok {
		self.Notify(pid, msg)
	}
}

func (self *Table) SendActionCall(pos seat) {

	// 叫分结束
	if len(self.deck_info.call_info) >= int(seat_max) {
		self.CalcCall()
		return
	}

	self.call_pos = pos

	self.l.Info("%v 开始叫分", pos_to_string(pos))

	// 设置开始时间
	self.action_ts = time.Now()

	msg := &pb.CallScoreBroadcast{
		Pos: int32(pos),
	}

	for _, v := range self.deck_info.call_info {
		msg.History = append(msg.History, v.score)
	}

	// 广播发消息
	self.Broadcast(msg)
}

func (self *Table) CalcCall() {
	var lord seat
	var score int32

	for _, v := range self.deck_info.call_info {
		if v.score > score {
			lord = v.pos
			score = v.score
		}
	}

	self.deck_info.caca_info.lord = lord
	self.deck_info.caca_info.score = score

	// 打牌 or 流局
	if score > 0 {
		self.deck_info.caca_info.draw = false
		self.l.Info("叫分结算: %v 是地主，叫了 %v 分，底牌是: %v", pos_to_string(lord), score, poker.CardsToString(self.cards[51:]))
		self.Switch(stage_play)
	} else {
		self.deck_info.caca_info.draw = true
		self.l.Info("叫分结算: 流局")
		self.Switch(stage_calc)
	}

}

func (self *Table) IsVictory() bool {
	return self.seats[self.play_pos].IsVictory()
}

// 玩家出牌
func (self *Table) PlayHand(cards []poker.Card, ci *poker.CardsInfo) {
	pos := self.play_pos

	if ci.IsBomb() {
		self.bombs++
	}

	// 删除手牌
	self.seats[pos].RemoveCards(cards)

	self.play_idx++
	self.play_pass = 0
	self.play_cards = ci

	// 通知谁出了牌
	self.Broadcast(&pb.PlayResultBroadcast{
		Pos:   int32(self.play_pos),
		Cards: poker.CardsToInt32(cards),
	})

	self.l.Info("%v 出牌：%v", pos_to_string(pos), poker.CardsToString(cards))
	self.l.Info("%v 剩余的牌为: %v", pos_to_string(self.play_pos), poker.CardsToString(self.seats[self.play_pos].data.cards))

	self.deck_info.play_info = append(self.deck_info.play_info, &play_info_t{})

	// 判断胜负
	if self.IsVictory() {
		loop.Next(func() {
			self.Switch(stage_calc)
		})

		return
	}

	self.play_pos = next_seat(self.play_pos)

	// 通知下一家出牌
	self.Broadcast(&pb.PlayBroadcast{
		Pos:   int32(self.play_pos),
		First: self.play_idx == 0,
	})

	self.l.Info("该 %v 出牌了， 首出：%v", pos_to_string(self.play_pos), self.play_idx == 0)
}

func (self *Table) PlayPass() {
	self.play_idx++
	self.play_pass++

	if self.play_pass == 2 {
		self.play_idx = 0
		self.play_pass = 0
	}

	// 通知谁出了牌
	self.Broadcast(&pb.PlayResultBroadcast{
		Pos:   int32(self.play_pos),
		Cards: nil,
	})

	self.l.Info("%v PASS", pos_to_string(self.play_pos))
	self.l.Info("%v 剩余的牌为: %v", pos_to_string(self.play_pos), poker.CardsToString(self.seats[self.play_pos].data.cards))

	self.play_pos = next_seat(self.play_pos)

	// 通知下一家出牌
	self.Broadcast(&pb.PlayBroadcast{
		Pos:   int32(self.play_pos),
		First: self.play_idx == 0,
	})

	self.l.Info("该 %v 出牌了， 首出：%v", pos_to_string(self.play_pos), self.play_idx == 0)
}
