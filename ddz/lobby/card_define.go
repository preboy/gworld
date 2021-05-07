package lobby

import (
	"fmt"
	"math/rand"
)

const (
	CardColor_Heart   int32 = 1 // 红
	CardColor_Spade   int32 = 2 // 黑
	CardColor_Diamond int32 = 3 // 方
	CardColor_Club    int32 = 4 // 梅
)

const (
	CardPoint_3 int32 = 3
	CardPoint_4 int32 = 4
	CardPoint_5 int32 = 5
	CardPoint_6 int32 = 6
	CardPoint_7 int32 = 7
	CardPoint_8 int32 = 8
	CardPoint_9 int32 = 9
	CardPoint_T int32 = 10
	CardPoint_J int32 = 11
	CardPoint_Q int32 = 12
	CardPoint_K int32 = 13
	CardPoint_A int32 = 14
	CardPoint_2 int32 = 16
)

const (
	CardValue_J1 int32 = 18000
	CardValue_J2 int32 = 20000
)

type Card int32

func NewPoker() (cards []Card) {
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

	// sort:  7 times equ chaos
	for x := 0; x < 7; x++ {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return
}

// ----------------------------------------------------------------------------
// member

func NewCardFromValue(v int32) Card {
	return Card(v)
}

func NewCard(color int32, point int32) Card {
	if color < CardColor_Heart || color > CardColor_Club {
		panic("NewCard:	invalid color")
	}

	if color < CardPoint_3 || color > CardPoint_2 || color == 15 {
		panic("NewCard:	invalid point")
	}

	return NewCardFromValue(point*1000 + color)
}

func (c Card) Color() int32 {
	return c.Value() % 1000
}

func (c Card) Point() int32 {
	return c.Value() / 1000
}

func (c Card) Value() int32 {
	return int32(c)
}

func (c Card) Valid() bool {
	v := c.Value()

	if v == CardValue_J1 || v == CardValue_J2 {
		return true
	}

	t := c.Color()
	p := c.Point()

	if t < CardColor_Heart || t > CardColor_Club {
		return false
	}

	if p < CardPoint_3 || p > CardPoint_2 || p == 15 {
		return false
	}

	return true
}

func (c Card) ToString() string {
	if c.Value() == CardValue_J1 {
		return "Joker1"
	}

	if c.Value() == CardValue_J2 {
		return "Joker2"
	}

	return fmt.Sprintf("%s%s", stringify_type(c.Color()), stringify_point(c.Point()))
}

// ----------------------------------------------------------------------------
// local

func stringify_type(v int32) string {
	switch v {
	case CardColor_Heart:
		return "♥"
	case CardColor_Spade:
		return "♠"
	case CardColor_Diamond:
		return "♦"
	case CardColor_Club:
		return "♣"
	default:
		panic(fmt.Sprintf("errType=%d", v))
	}
}

func stringify_point(v int32) string {
	switch v {
	case CardPoint_3:
		return "3"
	case CardPoint_4:
		return "4"
	case CardPoint_5:
		return "5"
	case CardPoint_6:
		return "6"
	case CardPoint_7:
		return "7"
	case CardPoint_8:
		return "8"
	case CardPoint_9:
		return "9"
	case CardPoint_T:
		return "10"
	case CardPoint_J:
		return "J"
	case CardPoint_Q:
		return "Q"
	case CardPoint_K:
		return "K"
	case CardPoint_A:
		return "A"
	case CardPoint_2:
		return "2"
	default:
		panic(fmt.Sprintf("errPoint=%d", v))
	}
}
