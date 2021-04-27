package lobby

func (self *player_data) AddCards(cards []Card) {
	self.data.cards = append(self.data.cards, cards...)
}
