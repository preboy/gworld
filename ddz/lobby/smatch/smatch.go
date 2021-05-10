package smatch

import (
	"gworld/core/utils"
	"gworld/ddz/comp"
)

const (
	MatchStage_Prepare int32 = iota + 0
	MatchStage_Running
	MatchStage_Completed
)

type SMatchConf struct {
	total_deck    int32
	match_name    string
	gamblers_name []string
}

type SMatch struct {
	mid   uint32
	conf  *SMatchConf
	table *Table
	stage int32
}

func NewSMatch(conf *SMatchConf) *SMatch {
	return &SMatch{
		mid:   utils.SeqU32(),
		conf:  conf,
		table: &Table{},
	}
}

// ----------------------------------------------------------------------------
// Impl for comp.IMatch

func (self *SMatch) GetMID() uint32 {
	return self.mid
}

func (self *SMatch) GetName() string {
	return self.conf.match_name
}

func (self *SMatch) IsOver() bool {
	return self.stage == MatchStage_Completed
}

func (self *SMatch) OnUpdate() {
	self.table.OnUpdate()
}

func (self *SMatch) OnMessage(pid string, req comp.IMessage, res comp.IMessage) {
	switch req.GetOP() {

	case 1: // join

	default:
		self.table.OnMessage(pid, req, res)
	}
}
