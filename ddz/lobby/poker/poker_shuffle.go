package poker

import (
	"math/rand"
)

// 纯随机洗牌
func shuffle_random(cards []Card) {
	// sort:  7 times equ chaos
	for x := 0; x < 7; x++ {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}
}
