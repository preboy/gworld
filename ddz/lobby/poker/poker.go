package poker

import "math/rand"

// ----------------------------------------------------------------------------
// export

func CreatePoker() (cards []Card) {
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

	cards_shuffle(cards)

	return
}

// ----------------------------------------------------------------------------
// local

func int_random(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func cards_random(cards []Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

// 纯随机洗牌
func cards_shuffle(cards []Card) {
	// sort:  7 times equ chaos
	for x := 0; x < 7; x++ {
		cards_random(cards)
	}
}
