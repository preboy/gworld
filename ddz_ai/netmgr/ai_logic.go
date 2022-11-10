package netmgr

import (
	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby/poker"
	"gworld/ddz/pb"
	"strconv"
)

type seat_t int32

const (
	seat_lord      seat_t = 0 + iota // 地主位
	seat_lord_next                   // 地主下一家
	seat_lord_prev                   // 地主上一家
)

var (
	_ai = &AILogic{}
)

// 方位数据(对自己来说没意义)
type pos_data_t struct {
	left_count int32        // 剩余牌数量
	cards_poss []poker.Card // 可能的牌 	自己(空)
	cards_deal []poker.Card // 出过的牌
}

// 一手
type hand_t struct {
	pos   int32
	cards []poker.Card
}

// 一回合
type round_t struct {
	hands []*hand_t
}

type AILogic struct {
	c *connector

	pos_me   int32 // 我的位置
	pos_lord int32 // 地主位置

	seat seat_t // 相对于地主，我的顺位

	cards        []poker.Card // 我手上的牌
	cards_left   []poker.Card // 其它人手上的牌之和(除去已被打出的)
	cards_bottom []poker.Card // 底牌(归地主所有)

	plrs map[int32]*pos_data_t

	rounds []*round_t // 本副出牌记录
}

// ----------------------------------------------------------------------------
// 打牌逻辑

func (ai *AILogic) DealCard(pos int32, arr []int32) {

	cards, valid := poker.CardsFromInt32(arr)
	if !valid {
		log.Info("fuck server")
		block.Done()
		return
	}

	poker.CardsSort(cards)

	ai.pos_me = pos
	ai.cards = cards

	log.Info("一副开始了: pos=%v, cards=%v", pos_to_string(pos), poker.CardsToString(cards))

	ai.on_init()
}

func (ai *AILogic) CallScoreBroadcast(pos int32, score []int32) {
	log.Info("该 %v 叫分了, %v", pos_to_string(pos), score)

	if pos == ai.pos_me {
		msg := &pb.CallScoreRequest{
			Score: ai.ai_call(),
		}
		ai.SendMessage(msg)
	}
}

func (ai *AILogic) CallScoreResponse(err int32) {
	if err == gconst.Err_OK {
		log.Info("叫分 OK")
	} else {
		log.Info("叫分 Error")
	}
}

func (ai *AILogic) CallScoreResultBroadcast(pos int32, score int32) {
	log.Info("%v 叫了 %v 分", pos_to_string(pos), score)
}

func (ai *AILogic) CallScoreCalcBroadcast(draw bool, lord int32, score int32, arr []int32) {
	if draw {
		log.Info("流局")
		return
	}

	cards, _ := poker.CardsFromInt32(arr)
	log.Info("%v 是地主，叫了 %v 分, 底牌: %v", pos_to_string(lord), score, poker.CardsToString(cards))

	ai.pos_lord = lord
	ai.calc_seat()

	ai.cards_bottom = cards

	// lord is me
	if ai.is_lord() {
		ai.cards = append(ai.cards, cards...)
		poker.CardsSort(ai.cards)
		log.Info("%v 是我<地主>，我的最终牌: %v", pos_to_string(ai.pos_me), poker.CardsToString(ai.cards))
	}

	ai.on_call_calc()
}

func (ai *AILogic) PlayBroadcast(pos int32, first bool) {
	log.Info("该 %v 出牌了，首出：%v", pos_to_string(pos), first)

	if pos != ai.pos_me {
		return
	}

	cards := ai.ai_play(first)

	msg := &pb.PlayRequest{
		Cards: poker.CardsToInt32(cards),
	}

	log.Info("[我]出牌: %v", poker.CardsToString(cards))

	ai.SendMessage(msg)
}

func (ai *AILogic) PlayResponse(err_code int32) {
	if err_code == gconst.Err_OK {
		log.Info("出牌返回 OK")
	} else {
		log.Info("出牌返回 ERR")
	}
}

func (ai *AILogic) PlayResultBroadcast(pos int32, first bool, arr []int32) {
	cards, _ := poker.CardsFromInt32(arr)
	log.Info("%v 出牌 ：%v", pos_to_string(pos), poker.CardsToString(cards))

	ai.on_play(pos, first, cards)

	if pos == ai.pos_me {
		log.Info("[我]剩下的牌为: %v", poker.CardsToString(ai.cards))
	}
}

func (ai *AILogic) DeckEndBroadcast(score []int32) {
	log.Info("结算: %v", score)
}

// ----------------------------------------------------------------------------
// local member

func (ai *AILogic) on_init() {
	ai.plrs = map[int32]*pos_data_t{}

	for i := int32(0); i < 3; i++ {
		ai.plrs[i] = &pos_data_t{
			left_count: 17,
		}
	}

	ai.cards_left, _ = poker.CardsRemove(poker.NewPoker(), ai.cards)
}

func (ai *AILogic) on_call_calc() {
	ai.plrs[ai.pos_lord].left_count += 3

	if ai.is_lord() {
		ai.cards_left, _ = poker.CardsRemove(ai.cards_left, ai.cards_bottom)
	}
}

func (ai *AILogic) on_play(pos int32, first bool, cards []poker.Card) {
	if first {
		ai.rounds = append(ai.rounds, &round_t{})
	}

	ai.add_play(pos, cards)

	if ai.pos_me == pos {
		ai.cards, _ = poker.CardsRemove(ai.cards, cards)
	} else {
		ai.cards_left, _ = poker.CardsRemove(ai.cards_left, cards)
	}

	if len(cards) != 0 {
		ai.plrs[pos].cards_deal = append(ai.plrs[pos].cards_deal, cards...)
		ai.plrs[pos].left_count -= int32(len(cards))
	}
}

func (ai *AILogic) is_lord() bool {
	return ai.pos_me == ai.pos_lord
}

func (ai *AILogic) get_mate_pos() int32 {
	if ai.is_lord() {
		panic("is lord")
	}

	for i := int32(0); i < 3; i++ {
		if i == ai.pos_lord || i == ai.pos_me {
			continue
		}

		return i
	}

	panic("not found mate")
}

// 计算我的座次信息
func (ai *AILogic) calc_seat() {
	if ai.is_lord() {
		ai.seat = seat_lord
		return
	}

	if (ai.pos_lord+1)%3 == ai.pos_me {
		ai.seat = seat_lord_next
	} else {
		ai.seat = seat_lord_prev
	}
}

func (ai *AILogic) add_play(pos int32, cards []poker.Card) {
	l := len(ai.rounds)
	r := ai.rounds[l-1]
	r.hands = append(r.hands, &hand_t{pos, cards})
}

func (ai *AILogic) prev_play() []poker.Card {
	l := len(ai.rounds)
	if l == 0 {
		return nil
	}

	r := ai.rounds[l-1]

	for i := len(r.hands) - 1; i >= 0; i-- {
		if len(r.hands[i].cards) != 0 {
			return r.hands[i].cards
		}
	}

	return nil
}

// ----------------------------------------------------------------------------
// other

func (ai *AILogic) SetConnector(c *connector) {
	ai.c = c
}

func (ai *AILogic) SendMessage(msg comp.IMessage) {
	if ai.c == nil {
		return
	}

	ai.c.SendMessage(msg)
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
