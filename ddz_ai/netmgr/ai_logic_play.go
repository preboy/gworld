package netmgr

import "gworld/ddz/lobby/poker"

type data struct {
	left_count int32        // 剩余牌数量
	left_cards []poker.Card // 自己(空)  其它人（可能的牌）
	deal_cards []poker.Card // 出过的牌
}

type hand struct {
	pos   int32
	cards []poker.Card
}

type round struct {
	hands []*hand
}

// ----------------------------------------------------------------------------
// event

func (self *AILogic) on_init() {
	self.plrs = map[int32]*data{}

	for i := int32(0); i < 3; i++ {
		self.plrs[i] = &data{
			left_count: 17,
		}
	}

	self.left_cards, _ = poker.CardsRemove(poker.NewPoker(), self.cards)
}

func (self *AILogic) on_calc() {
	self.plrs[self.lord_pos].left_count += 3

	if self.lord_pos == self.pos {
		self.left_cards, _ = poker.CardsRemove(self.left_cards, self.lord_cards)
	}
}

func (self *AILogic) on_play(pos int32, first bool, cards []poker.Card) {
	if first {
		self.rounds = append(self.rounds, &round{})
	}

	self.add_play(pos, cards)

	if self.pos == pos {
		self.cards, _ = poker.CardsRemove(self.cards, cards)
		return
	}

	if len(cards) != 0 {
		self.left_cards, _ = poker.CardsRemove(self.left_cards, cards)
		self.plrs[pos].deal_cards = append(self.plrs[pos].deal_cards, cards...)
		self.plrs[pos].left_count -= int32(len(cards))
	}
}

func (self *AILogic) play(first bool) (cards []poker.Card, ok bool) {

	a1 := poker.NewAnalyse(self.cards)
	a2 := poker.NewAnalyse(self.left_cards)

	if first {

	} else {
		cards_prev := self.prev_play()
		if cards_prev == nil {
			panic("cards_prev == nil")
		}

		// 上一首不为Nil的牌
		ci := poker.CardsAnalyse(cards_prev)

		// 找到大过他的牌
		a1.Exceed(ci)

	}

	return
}

// ----------------------------------------------------------------------------
// local

func (self *AILogic) is_lord() bool {
	return self.pos == self.lord_pos
}

func (self *AILogic) get_friend_pos() int32 {
	if ai.is_lord() {
		panic("is lord")
	}

	for i := int32(0); i < 3; i++ {
		if i == self.lord_pos || i == self.pos {
			continue
		}

		return i
	}

	panic("not found friend")
}

func (self *AILogic) add_play(pos int32, cards []poker.Card) {
	l := len(self.rounds)
	r := self.rounds[l-1]
	r.hands = append(r.hands, &hand{pos, cards})
}

func (self *AILogic) prev_play() []poker.Card {
	l := len(self.rounds)
	r := self.rounds[l-1]

	for i := len(r.hands) - 1; i >= 0; i-- {
		if len(r.hands[i].cards) != 0 {
			return r.hands[i].cards
		}
	}

	return nil
}
