package smatch

import (
	"gworld/ddz/comp"
	"gworld/ddz/lobby/poker"
)

// 一个人本副牌的信息
type deck_data struct {
	cards []poker.Card
}

type gambler_table_t struct {
	m *Table

	pid  string
	pos  seat
	data *deck_data
}

// ----------------------------------------------------------------------------

func (self *gambler_table_t) AddCards(cards []poker.Card) {
	self.data.cards = append(self.data.cards, cards...)
	poker.CardsSort(self.data.cards)
}

// 获取最小的牌
func (self *gambler_table_t) GetDefaultCards() (ret []poker.Card) {
	poker.CardsSort(self.data.cards)

	l := len(self.data.cards)
	ret = self.data.cards[l-1 : l]

	return
}

// 是否手上有这些牌
func (self *gambler_table_t) ExistCards(cards []poker.Card) bool {
	if len(cards) == 0 {
		panic("nil cards")
	}

	for _, c := range cards {
		if !poker.CardsExist(self.data.cards, c) {
			return false
		}
	}

	return true
}

func (self *gambler_table_t) RemoveCards(cards []poker.Card) {
	if len(cards) == 0 {
		return
	}

	new_cards, ok := poker.CardsRemove(self.data.cards, cards)

	if ok {
		poker.CardsSort(new_cards)
		self.data.cards = new_cards
	}
}

func (self *gambler_table_t) IsVictory() bool {
	return len(self.data.cards) == 0
}

func (self *gambler_table_t) SendMessage(msg comp.IMessage) {
	gbr := comp.GM.FindGambler(self.pid)
	if gbr != nil {
		gbr.SendMessage(msg)
	}
}
