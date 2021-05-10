package poker

func new_poker() (cards []Card) {
	// 3 ~ A * 4
	for p := CardPoint_3; p <= CardPoint_A; p++ {
		for c := CardColor_Heart; c <= CardColor_Club; c++ {
			cards = append(cards, NewCard(c, p))
		}
	}
	// 2 * 4
	for c := CardColor_Heart; c <= CardColor_Club; c++ {
		cards = append(cards, NewCard(c, CardPoint_2))
	}

	// Joker1 Joker2
	cards = append(cards, NewCardFromValue(CardValue_J1))
	cards = append(cards, NewCardFromValue(CardValue_J2))

	return
}

// ----------------------------------------------------------------------------
// export

func CreatePoker() []Card {
	cards := new_poker()
	shuffle_random(cards)
	return cards
}
