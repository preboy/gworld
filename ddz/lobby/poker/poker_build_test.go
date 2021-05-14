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

	b := NewPokerBuild()
	b.Build_CardsTypeA_SEQ(3, 7)

	cards := b.Done()

	fmt.Println(CardsToString(cards[:17]))
	fmt.Println(CardsToString(cards[17:34]))
	fmt.Println(CardsToString(cards[34:51]))
	fmt.Println(CardsToString(cards[51:]))
}
