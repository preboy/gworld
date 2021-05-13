package smatch

import (
	"gworld/core/utils"
	"gworld/ddz/comp"
)

type (
	match_stage = int32
)

const (
	MatchStage_Prepare int32 = iota + 0
	MatchStage_Running
	MatchStage_Completed
)

type gambler_match_t struct {
	score         int // 总分
	win_count     int // 胜次数
	lost_count    int // 败次数
	load_count    int // 地主次数
	peasant_count int // 农民次数
}

type SMatchConf struct {
	TotalDeck    int32
	MatchMame    string
	GamblerNames []string
}

type SMatch struct {
	mid      uint32
	conf     *SMatchConf
	stage    match_stage
	table    *Table
	gamblers map[string]*gambler_match_t
}

func NewSMatch(conf *SMatchConf) *SMatch {
	m := &SMatch{
		mid:      utils.SeqU32(),
		conf:     conf,
		stage:    MatchStage_Prepare,
		table:    &Table{},
		gamblers: map[string]*gambler_match_t{},
	}

	m.table.m = m

	return m
}

// ----------------------------------------------------------------------------
// Impl for comp.IMatch

func (self *SMatch) GetMID() uint32 {
	return self.mid
}

func (self *SMatch) GetName() string {
	return self.conf.MatchMame
}

func (self *SMatch) IsOver() bool {
	return self.stage == MatchStage_Completed
}

func (self *SMatch) OnUpdate() {
	self.table.OnUpdate()
}

func (self *SMatch) OnMessage(pid string, req comp.IMessage, res comp.IMessage) {
	switch req.GetOP() {
	default:
		self.table.OnMessage(pid, req, res)
	}
}

func (self *SMatch) Sit(pid string) bool {
	if self.stage != MatchStage_Prepare {
		return false
	}

	// 名字是否允许
	plr := comp.GM.FindGambler(pid)
	if plr == nil {
		return false
	}

	find := false
	name := plr.GetName()
	for _, v := range self.conf.GamblerNames {
		if v == name {
			find = true
		}
	}

	if !find {
		return false
	}

	if self.table.Sit(pid) {
		self.gamblers[pid] = &gambler_match_t{}
	}

	return true
}
