package lobby

import (
	"gworld/core/utils"
	"time"
)

type (
	SEAT  int
	STAGE int
)

const (
	SEAT_EAST  SEAT = 0
	SEAT_SOUTH SEAT = 1
	SEAT_WEST  SEAT = 2
	SEAT_MAX   SEAT = 3
)

const (
	stage_prepare STAGE = iota + 0 // 未开始
	stage_deal                     // 发牌
	stage_call                     // 叫分
	stage_play                     // 出牌
	stage_calc                     // 结算
	stage_over                     // 已结束
	stage_MAX
)

// 一个人本副牌的信息
type deck_data struct {
	cards []Card
}

type player_data struct {
	m *Match

	pid  string
	pos  SEAT
	data *deck_data

	// stat
	score_total   int // 总分
	win_count     int // 胜次数
	lost_count    int // 败次数
	load_count    int // 地主次数
	peasant_count int // 农民次数
}

type Match struct {
	ID   uint32
	pids []string

	seats [SEAT_MAX]*player_data // 3个方位的pid
	stage STAGE

	deck_index int // 当前牌副数
	deck_total int // 总牌副数

	host_pos   SEAT // 首叫方位
	call_pos   SEAT // 叫分方位
	call_score int  // 最高叫分

	action_ts time.Time

	deck_data         *deck_info_t   // 当前副数据
	deck_data_history []*deck_info_t // 历史数据
}

func NewMatch() *Match {
	return &Match{
		ID: utils.SeqU32(),
	}
}

func (self *Match) Init(pids []string) {
	self.pids = pids
	self.stage = stage_prepare

	self.deck_index = 0
	self.deck_total = 10

	self.host_pos = SEAT_EAST

	for i := SEAT_EAST; i < SEAT_MAX; i++ {
		self.seats[i] = &player_data{
			m:    self,
			pid:  pids[i],
			pos:  i,
			data: &deck_data{},
		}
	}

	FSM[self.stage].OnEnter(self)
}

func (self *Match) OnUpdate() {
	FSM[self.stage].OnUpdate(self)
}

func (self *Match) Switch(stage STAGE) {
	FSM[self.stage].OnLeave(self)
	self.stage = stage
	FSM[self.stage].OnEnter(self)
}

func (self *Match) Over() bool {
	return self.stage == stage_over
}

func (self *Match) Exist(pid string) bool {
	for _, v := range self.pids {
		if v == pid {
			return true
		}
	}

	return false
}

func (self *Match) DeckOpen() {
	self.deck_index++

	for i := SEAT_EAST; i < SEAT_MAX; i++ {
		self.seats[i].data = &deck_data{}
	}

	self.deck_data = &deck_info_t{
		index: self.deck_index,
		start: time.Now().Unix(),
		deal:  nil,
		call:  nil,
		callr: &call_result_t{},
		play:  nil,
		calc:  &calc_info_t{},
	}

	self.call_pos = self.host_pos
	self.call_score = 0
}

func (self *Match) DeckClosed() {
	self.host_pos = next_seat(self.host_pos)
	self.deck_data_history = append(self.deck_data_history, self.deck_data)
}

func (self *Match) NextDeck() {
	if self.deck_index < self.deck_total {
		self.Switch(stage_deal)
	} else {
		self.Switch(stage_over)
	}
}

func (self *Match) Broadcast() {

}

func (self *Match) SetActionCall(cp SEAT) {
	self.call_pos = cp

	// 设置开始时间
	self.action_ts = time.Now()

	// 广播发消息
	self.seats[self.call_pos].SetActionCall()
}
