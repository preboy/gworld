package lobby

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

	cards, ok := cards_from_int32(arr)
	if !ok {
		t.Fatal("invalid card")
	}

	ci := get_cards_info(cards)

	fmt.Println(ci.ToString())
}
