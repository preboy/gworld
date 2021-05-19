package poker

import (
	"fmt"
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

	CardPoint_J1 int32 = 18
	CardPoint_J2 int32 = 20
)

const (
	CardValue_J1 int32 = 18000
	CardValue_J2 int32 = 20000
)

const (
	Card_PlaceHolder_A  int32 = -1
	Card_PlaceHolder_AA int32 = -2
)

// ----------------------------------------------------------------------------
// Card

type Card int32

func NewCardFromValue(v int32) Card {
	return Card(v)
}

func NewCard(color int32, point int32) Card {
	if color < CardColor_Heart || color > CardColor_Club {
		panic("NewCard:	invalid color")
	}

	if point < CardPoint_3 || point > CardPoint_2 || point == 15 {
		panic("NewCard:	invalid point")
	}

	return NewCardFromValue(point*1000 + color)
}

func NewCard_PlaceHolder_A() Card {
	return Card(Card_PlaceHolder_A)
}

func NewCard_PlaceHolder_AA() Card {
	return Card(Card_PlaceHolder_AA)
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

func (c Card) IsPlaceHolder() bool {
	v := c.Value()
	return v == Card_PlaceHolder_A || v == Card_PlaceHolder_AA
}

func (c Card) ToString() string {
	if c.Value() == CardValue_J1 {
		return "Joker1"
	}

	if c.Value() == CardValue_J2 {
		return "Joker2"
	}

	if c.IsPlaceHolder() {
		if c.Value() == Card_PlaceHolder_A {
			return "[_]"
		}
		if c.Value() == Card_PlaceHolder_AA {
			return "[__]"
		}
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
