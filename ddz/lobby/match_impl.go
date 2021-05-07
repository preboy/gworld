package lobby

func (self *player_data) AddCards(cards []Card) {
	self.data.cards = append(self.data.cards, cards...)
}

// 获取最小的牌
func (self *player_data) GetDefaultCards() (ret []Card) {
	cards_sort(self.data.cards)

	l := len(self.data.cards)
	ret = self.data.cards[l-1 : l]
	self.data.cards = self.data.cards[:l-1]

	return
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
		l, e := remove_card(local, NewCardFromValue(v))
		if !e {
			return false
		}

		local = l
	}

	return true
}

// ----------------------------------------------------------------------------
// local
func valid_cards(cards []Card) bool {

	return false
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