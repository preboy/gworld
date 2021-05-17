package poker

import "math/rand"

type BuildType = int32

const (
	BuildType_ABCDE BuildType = iota + 0
	BuildType_AAABBB
	BuildType_AABBCC
	BuildType_AAAA
)

type BuildParam struct {
	BType  BuildType
	Count  int32
	Length int32
}

// ----------------------------------------------------------------------------
// export

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

	return
}

func CreatePoker() []Card {
	cards := NewPoker()
	cards_shuffle(cards)
	return cards
}

func BuildPoker(params []*BuildParam) (cards []Card) {
	b := NewPokerBuild()

	for _, v := range params {
		switch v.BType {
		case BuildType_ABCDE:
			b.Build_ABCDE(v.Count, v.Length)
		case BuildType_AAABBB:
			b.Build_AAABBB(v.Count, v.Length)
		case BuildType_AABBCC:
			b.Build_AABBCC(v.Count, v.Length)
		case BuildType_AAAA:
			b.Build_AAAA(v.Count)
		}
	}

	return b.Finish()
}

// ----------------------------------------------------------------------------
// local

func int_random(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func cards_random(cards []Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

// 纯随机洗牌
func cards_shuffle(cards []Card) {
	// sort:  7 times equ chaos
	for x := 0; x < 7; x++ {
		cards_random(cards)
	}
}
