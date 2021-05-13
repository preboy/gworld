package netmgr

import (
	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby/poker"
	"gworld/ddz/pb"
	"math/rand"
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
	log.Info("该 %v 叫分了, %v", pos_to_string(pos), score)

	if pos == self.pos {
		msg := &pb.CallScoreRequest{
			Score: rand.Int31n(3),
		}
		self.c.SendMessage(msg)
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
	log.Info("%v 是地主，叫了 %v 分, %v", pos_to_string(lord), score, poker.CardsToString(cards))

	// me
	cards_lord, _ := poker.CardsFromInt32(arr)
	if lord == self.pos {
		self.cards = append(self.cards, cards_lord...)
	}
}

func (self *AILogic) PlayBroadcast(pos int32, first bool) {
	log.Info("该 %v 出牌了，首出：%v", pos_to_string(pos), first)

	if pos != self.pos {
		return
	}

	poker.CardsSort(self.cards)

	msg := &pb.PlayRequest{}

	cards := []poker.Card{}

	if first {
		l := len(self.cards)
		cards = self.cards[l-1:]
		self.cards = self.cards[:l-1]
	}

	msg.Cards = poker.CardsToInt32(cards)

	log.Info("出牌: %v", poker.CardsToString(cards))
	log.Info("剩下的牌为: %v", poker.CardsToString(self.cards))

	self.c.SendMessage(msg)
}

func (self *AILogic) PlayResponse(err_code int32) {
	if err_code == gconst.Err_OK {
		log.Info("出牌返回OK")
	} else {
		log.Info("出牌返回 ERR")
	}
}

func (self *AILogic) PlayResultBroadcast(pos int32, arr []int32) {
	cards, _ := poker.CardsFromInt32(arr)
	log.Info("%v 出牌 ：%v", pos_to_string(pos), poker.CardsToString(cards))
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
	case 3:
		return "<北>"
	default:
		panic("方位错误" + strconv.Itoa(int(pos)))
	}
}
