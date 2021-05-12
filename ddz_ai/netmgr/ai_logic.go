package netmgr

import (
	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/lobby/poker"
	"strconv"
)

var (
	ai = &AILogic{}
)

type AILogic struct {
	c *connector

	// base
	pos   int32
	cards []poker.Card
}

// ----------------------------------------------------------------------------
// member

func (self *AILogic) Init(c *connector, pos int32, arr []int32) {

	cards, valid := poker.CardsFromInt32(arr)
	if !valid {
		log.Info("fuck server")
		block.Signal()
		return
	}

	self.c = c
	self.pos = pos
	self.cards = cards

	log.Info("一副开始了: pos=%v, cards=%v", pos_to_string(pos), poker.CardsToString(cards))
}

func (self *AILogic) CallScoreBroadcast(pos int32, score []int32) {
	log.Info("该 %v 叫分了, %v %v", pos_to_string(pos), score)
	// me
}

func (self *AILogic) CallScoreResultBroadcast(pos int32, score int32) {
	//
	log.Info("%v 叫了 %v 分", pos_to_string(pos), score)
}

func (self *AILogic) CallScoreCalcBroadcast(draw bool, lord int32, score int32, arr []int32) {

	if draw {
		log.Info("流局")
	} else {
		cards, _ := poker.CardsFromInt32(arr)
		log.Info("%v 是地主，叫了 %v 分, %v", pos_to_string(lord), score, poker.CardsToString(cards))
	}
}

func (self *AILogic) PlayBroadcast(pos int32, first bool) {
	log.Info("该 %v 出牌了，首出：%v", pos_to_string(pos), first)
}

// ----------------------------------------------------------------------------
// local
func pos_to_string(pos int32) string {
	switch pos {
	case 0:
		return "<东>"
	case 1:
		return "<南>"
	case 2:
		return "<西>"
	case 3:
		return "<北>"
	default:
		panic("方位错误" + strconv.Itoa(int(pos)))
	}
}
