package netmgr

import (
	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby/poker"
	"gworld/ddz/pb"
	"strconv"
)

var (
	ai = &AILogic{}
)

type AILogic struct {
	c *connector

	pos        int32        // 我的位置
	lord_pos   int32        // 地主位置
	lord_cards []poker.Card // 底牌

	cards      []poker.Card // 我手上的牌
	left_cards []poker.Card // 其它人手上的的牌

	plrs map[int32]*data

	rounds []*round // 本副出牌记录
}

// ----------------------------------------------------------------------------
// member

func (self *AILogic) Init(c *connector, pos int32, arr []int32) {

	cards, valid := poker.CardsFromInt32(arr)
	if !valid {
		log.Info("fuck server")
		block.Done()
		return
	}

	poker.CardsSort(cards)

	self.c = c
	self.pos = pos
	self.cards = cards

	log.Info("一副开始了: pos=%v, cards=%v", pos_to_string(pos), poker.CardsToString(cards))

	self.on_init()
}

func (self *AILogic) CallScoreBroadcast(pos int32, score []int32) {
	log.Info("该 %v 叫分了, %v", pos_to_string(pos), score)

	if pos == self.pos {
		msg := &pb.CallScoreRequest{
			Score: self.ai_call(),
		}
		self.c.SendMessage(msg)
	}
}

func (self *AILogic) CallScoreResponse(err int32) {
	if err == gconst.Err_OK {
		log.Info("叫分OK")
	} else {
		log.Info("叫分Error")
	}
}

func (self *AILogic) CallScoreResultBroadcast(pos int32, score int32) {
	log.Info("%v 叫了 %v 分", pos_to_string(pos), score)
}

func (self *AILogic) CallScoreCalcBroadcast(draw bool, lord int32, score int32, arr []int32) {
	if draw {
		log.Info("流局")
		return
	}

	cards, _ := poker.CardsFromInt32(arr)
	log.Info("%v 是地主，叫了 %v 分, 底牌: %v", pos_to_string(lord), score, poker.CardsToString(cards))

	// lord is me
	if lord == self.pos {
		self.cards = append(self.cards, cards...)
		poker.CardsSort(self.cards)
		log.Info("%v 是我，我的最终牌: %v", pos_to_string(self.pos), poker.CardsToString(self.cards))
	}

	self.lord_pos = lord
	self.lord_cards = cards

	self.on_call_calc()
}

func (self *AILogic) PlayBroadcast(pos int32, first bool) {
	log.Info("该 %v 出牌了，首出：%v", pos_to_string(pos), first)

	if pos != self.pos {
		return
	}

	cards := self.ai_play(first)

	msg := &pb.PlayRequest{
		Cards: poker.CardsToInt32(cards),
	}

	log.Info("[我]出牌: %v", poker.CardsToString(cards))

	self.c.SendMessage(msg)
}

func (self *AILogic) PlayResponse(err_code int32) {
	if err_code == gconst.Err_OK {
		log.Info("出牌返回OK")
	} else {
		log.Info("出牌返回 ERR")
	}
}

func (self *AILogic) PlayResultBroadcast(pos int32, first bool, arr []int32) {
	cards, _ := poker.CardsFromInt32(arr)
	log.Info("%v 出牌 ：%v", pos_to_string(pos), poker.CardsToString(cards))

	self.on_play(pos, first, cards)

	if pos == self.pos {
		log.Info("[我]剩下的牌为: %v", poker.CardsToString(self.cards))
	}
}

func (self *AILogic) DeckEndBroadcast(score []int32) {
	log.Info("结算: %v", score)
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
	default:
		panic("方位错误" + strconv.Itoa(int(pos)))
	}
}
