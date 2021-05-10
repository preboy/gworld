package smatch

import "gworld/ddz/lobby/poker"

func (self *gambler) AddCards(cards []poker.Card) {
	self.data.cards = append(self.data.cards, cards...)
}

// 获取最小的牌
func (self *gambler) GetDefaultCards() (ret []poker.Card) {
	poker.CardsSort(self.data.cards)

	l := len(self.data.cards)
	ret = self.data.cards[l-1 : l]
	self.data.cards = self.data.cards[:l-1]

	return
}

// 是否手上有这些牌
func (self *gambler) ExistCards(cards []poker.Card) bool {
	if len(cards) == 0 {
		panic("nil cards")
	}

	local := []poker.Card{}
	for _, v := range self.data.cards {
		local = append(local, v)
	}

	for _, c := range cards {
		l, e := remove_card(local, c)
		if !e {
			return false
		}

		local = l
	}

	return true
}

func (self *gambler) RemoveCards(cards []poker.Card) {
	if len(cards) == 0 {
		return
	}

	new_cards := []poker.Card{}

	for _, c := range self.data.cards {
		if !exist_card(cards, c) {
			new_cards = append(new_cards, c)
		}
	}

	self.data.cards = new_cards
}

func (self *gambler) IsVictory() bool {
	return len(self.data.cards) == 0
}

// ----------------------------------------------------------------------------
// local

func exist_card(cards []poker.Card, c poker.Card) bool {
	for _, v := range cards {
		if v == c {
			return true
		}
	}
	return false
}

func remove_card(cards []poker.Card, card poker.Card) (left []poker.Card, e bool) {
	for k, v := range cards {
		if v == card {
			e = true
			left = append(left, cards[:k]...)
			left = append(left, cards[k+1:]...)
			return
		}
	}

	return cards, false
}
