package lobby

import (
	"fmt"
	"math/rand"
)

type CardsType int

const (
	CardsTypeNIL              CardsType = iota + 0 // 非法
	CardsTypeA                                     // 单牌
	CardsTypeAA                                    // 一对
	CardsTypeJJ                                    // 王炸
	CardsTypeAAA                                   // 三不带
	CardsTypeAAAX                                  // 三带单
	CardsTypeAAAXX                                 // 三带对
	CardsTypeAAAA                                  // 炸弹
	CardsTypeAAAAXY                                // 四带二
	CardsTypeAAAAXXYY                              // 四带二对
	CardsTypeABCDE                                 // 顺子(大于5)
	CardsTypeAABBCC                                // 顺对
	CardsTypeAAABBB                                // 飞机
	CardsTypeAAABBBXY                              // 飞机带单
	CardsTypeAAABBBXXYY                            // 飞机带对
	CardsTypeAAAABBBBXYZM                          // 飞机顺带单
	CardsTypeAAAABBBBXXYYZZMM                      // 飞机顺带对
)

const (
	CardType_Heart   = 1 // 红
	CardType_Spade   = 2 // 黑
	CardType_Diamond = 3 // 方
	CardType_Club    = 4 // 梅
)

const (
	CardPoint_3 = 3
	CardPoint_4 = 4
	CardPoint_5 = 5
	CardPoint_6 = 6
	CardPoint_7 = 7
	CardPoint_8 = 8
	CardPoint_9 = 9
	CardPoint_T = 10
	CardPoint_J = 11
	CardPoint_Q = 12
	CardPoint_K = 13
	CardPoint_A = 14
	CardPoint_2 = 16

	CardPoint_J1 = 18000
	CardPoint_J2 = 20000
)

type Card int

func NewPoker() (cards []Card) {
	// 3 ~ A * 4
	for p := CardPoint_3; p <= CardPoint_A; p++ {
		for t := CardType_Heart; t <= CardType_Club; t++ {
			v := p*1000 + t
			cards = append(cards, Card(v))
		}
	}
	// 2 * 4
	for t := CardType_Heart; t <= CardType_Club; t++ {
		v := CardPoint_2*1000 + t
		cards = append(cards, Card(v))
	}

	// Joker1 Joker2
	cards = append(cards, Card(CardPoint_J1))
	cards = append(cards, Card(CardPoint_J2))

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

func (c Card) Type() int {
	return int(c) % 1000
}

func (c Card) Point() int {
	return int(c) / 1000
}

func (c Card) Valid() bool {
	if c == CardPoint_J1 || c == CardPoint_J2 {
		return true
	}

	t := c.Type()
	p := c.Point()

	if t < CardType_Heart || t > CardType_Club {
		return false
	}

	if p < CardPoint_3 || p > CardPoint_2 || p == 15 {
		return false
	}

	return true
}

func (c Card) ToString() string {
	if c == Card(CardPoint_J1) {
		return "Joker1"
	}

	if c == Card(CardPoint_J2) {
		return "Joker2"
	}

	return fmt.Sprintf("%s%s", stringify_type(c.Type()), stringify_point(c.Point()))
}

// ----------------------------------------------------------------------------
// local

func stringify_type(v int) string {
	switch v {
	case CardType_Heart:
		return "♥"
	case CardType_Spade:
		return "♠"
	case CardType_Diamond:
		return "♦"
	case CardType_Club:
		return "♣"
	default:
		panic(fmt.Sprintf("errType=%d", v))
	}
}

func stringify_point(v int) string {
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
