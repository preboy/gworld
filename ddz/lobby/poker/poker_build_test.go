package poker

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestA45(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		fmt.Println("result:", random_angle_45(3, 7))
	}
}

func TestPokerBuild(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		b := NewPokerBuild()
		b.Build_ABCDE(3, 7)
		b.Build_AAABBB(3, 4)
		b.Build_AABBCC(3, 8)
		b.Build_AAAA(3)

		cards := b.Finish()

		if len(cards) != 54 {
			t.Fail()
		}

		fmt.Println("------------------------------------------")
		fmt.Println(CardsToString(cards[:17]))
		fmt.Println(CardsToString(cards[17:34]))
		fmt.Println(CardsToString(cards[34:51]))
		fmt.Println(CardsToString(cards[51:]))
	}
}
