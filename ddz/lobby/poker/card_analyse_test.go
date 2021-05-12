package poker

import (
	"fmt"
	"testing"
)

func TestAnalyse(t *testing.T) {
	arr := []int32{
		3001,
		3001,
		4001,
		4001,
		5001,
		5001,
	}

	cards, ok := CardsFromInt32(arr)
	if !ok {
		t.Fatal("invalid card")
	}

	ci := CardsAnalyse(cards)

	fmt.Println(ci.ToString())
}
