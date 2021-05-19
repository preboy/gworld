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

func TestExceed(t *testing.T) {

	// 目标牌
	arr := []int32{
		3001,
		3002,
		3003,
		4001,
		4002,
	}

	// 手牌
	hand_arr := []int32{
		4003,
		4004,
		5003,
		5004,
		6003,
		6004,
		7003,
		7004,
		8002,
		8003,
		8004,

		9001, 9002, 9003, 9004,
		18000, 20000,
	}

	cards, ok := CardsFromInt32(arr)
	if !ok {
		t.Fatal("invalid arr")
	}

	ci := CardsAnalyse(cards)

	hand_cards, ok := CardsFromInt32(hand_arr)
	if !ok {
		t.Fatal("invalid hand_arr")
	}

	a := NewAnalyse(hand_cards)

	fmt.Println("dst:", CardsToString(cards), arr)

	for i, v := range a.Exceed(ci) {
		fmt.Println("the <", i, "> groups: ", CardsToString(v), v)
	}
}
