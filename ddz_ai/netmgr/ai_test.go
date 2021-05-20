package netmgr

import (
	"fmt"
	"gworld/ddz/lobby/poker"
	"testing"
)

func TestAI(t *testing.T) {
	arr := []int32{
		3001,
		3002,
		3003,
		3004,

		4001, 4002, 4003,

		5001, 5002,

		6001,

		7001, 8001, 9001, 10001, 11001,
	}

	cards, ok := poker.CardsFromInt32(arr)
	if !ok {
		panic("invalid arr")
	}

	c := cards_divide_abdef(cards)

	fmt.Println("class:", c, poker.CardsToString(cards))
}
