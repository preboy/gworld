package lobby

import (
	"gworld/core/utils"
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
	card []int

	bomb int // 炸弹数
}

type match_data struct {
	pid  string
	pos  SEAT
	data *deck_data
}

type Match struct {
	ID   uint32
	pids []string

	seats [SEAT_MAX]*match_data // 3个方位的pid
	stage STAGE

	deck_curr  int // 当前牌副数
	deck_total int // 总牌副数

	first_call SEAT // 首叫方位
}

func NewMatch() *Match {
	return &Match{
		ID: utils.SeqU32(),
	}
}

func (self *Match) Init(pids []string) {
	self.pids = pids
	self.stage = stage_prepare

	self.deck_curr = 0
	self.deck_total = 10

	self.first_call = SEAT_EAST

	for i := SEAT_EAST; i < SEAT_MAX; i++ {
		self.seats[i] = &match_data{
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

func (self *Match) InitDeck() {
	self.deck_curr++

	for i := SEAT_EAST; i < SEAT_MAX; i++ {
		self.seats[i].data = &deck_data{}
	}
}

func (self *Match) NextDeck() {
	if self.deck_curr < self.deck_total {
		self.Switch(stage_deal)
	} else {
		self.Switch(stage_over)
	}
}
