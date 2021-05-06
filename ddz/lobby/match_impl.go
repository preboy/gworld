package lobby

import "sort"

func (self *player_data) AddCards(cards []Card) {
	self.data.cards = append(self.data.cards, cards...)
}

// 获取最小的牌
func (self *player_data) GetDefaultCards() (ret []Card) {
	sort.Slice(self.data.cards, func(i, j int) bool {
		return self.data.cards[i] > self.data.cards[j]
	})

	l := len(self.data.cards)
	ret = self.data.cards[l-1 : l]
	self.data.cards = self.data.cards[:l-1]

	return
}

func remove_card(cards []Card, card Card) (left []Card, e bool) {
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

// 是否手上有这些牌
func (self *player_data) ExistCards(cards []int32) bool {
	if len(cards) == 0 {
		panic("nil cards")
	}

	local := []Card{}
	for _, v := range self.data.cards {
		local = append(local, v)
	}

	for _, v := range cards {
		l, e := remove_card(local, Card(v))
		if !e {
			return false
		}

		local = l
	}

	return true
}
