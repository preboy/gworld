package poker

import (
	"math/rand"
)

const (
	cards_max_number int32 = 17
)

type card_wall map[int][]Card

func (self card_wall) init() {
	self = map[int][]Card{}
}

func (self card_wall) exist_seq(p2 int, length int, cnt int) bool {
	for i := p2 - length + 1; i <= length; i++ {
		if len(self[i]) < cnt {
			return false
		}
	}
	return true
}

func (w card_wall) expect_a_seq(length int) (ret []Card) {
	p1 := 3 + length - 1
	p2 := 14

	set := []int{}

	// seal
	for i := p1; i <= p2; i++ {
		if w.exist_seq(p1, length, 1) {
			set = append(set, i)
		}
	}

	if len(set) == 0 {
		return
	}

	r := rand.Intn(len(set))
	v := set[r]

	for i := v - length + 1; i <= length; i++ {
		var c Card
		w[i], c = cards_get(w[i])
		ret = append(ret, c)
	}

	return
}

type poker_build struct {
	wall card_wall

	draw [][]Card
}

// ----------------------------------------------------------------------------
// member

func NewPokerBuild() *poker_build {
	pb := &poker_build{}
	pb.init()

	return pb
}

func (self *poker_build) init() {
	self.wall.init()

	// 3 ~ A * 4
	for p := CardPoint_3; p <= CardPoint_A; p++ {
		var cards []Card
		for c := CardColor_Heart; c <= CardColor_Club; c++ {
			cards = append(cards, NewCard(c, p))
		}
		cards_random(cards)
		self.wall[p] = cards
	}

	// 2 * 4
	{
		var cards []Card
		for c := CardColor_Heart; c <= CardColor_Club; c++ {
			cards = append(cards, NewCard(c, CardPoint_2))
		}
		cards_random(cards)
		self.wall[CardPoint_2] = cards
	}

	// joker
	{
		c := NewCardFromValue(CardValue_J1)
		p := c.Point()
		self.wall[p] = append(self.wall[p], c)
	}

	// joker
	{
		c := NewCardFromValue(CardValue_J2)
		p := c.Point()
		self.wall[p] = append(self.wall[p], c)
	}
}

func (self *poker_build) Build() (ret []Card) {
	seats := [3][]Card{}

	// gen -> seat

	// wall -> seat

	for i := 0; i < 3; i++ {
		ret = append(ret, seats[i]...)
	}

	return
}

func (self *poker_build) Build_CardsTypeA_SEQ(cnt int32, len int32) bool {
	// args checking
	{
		if cnt <= 0 || cnt > 5 {
			return false
		}

		if len < 5 || len > 12 {
			return false
		}
	}

	for i := int32(0); i < cnt; i++ {
		l := random_angle_45(5, len)

		// wall 能否找到l连的
		self.wall.exist()

	}

	return true
}

// ----------------------------------------------------------------------------
// local

func random_angle_45(p1, p2 int32) int32 {
	if p1 == p2 {
		return p1
	}

	if p1 > p2 {
		panic("err args")
	}

	dt := p2 - p1 + 1

	total := int32(0)
	for i := int32(1); i <= dt; i++ {
		total += i
	}

	r := rand.Int31n(total) + 1
	sum := int32(0)

	for i := int32(1); i <= dt; i++ {
		sum += i

		if r <= sum {
			return p1 + dt - i
		}
	}

	panic("out of range")
}

func cards_get(cards []Card) (left []Card, c Card) {
	l := len(cards)

	c = cards[l-1]
	left = cards[:l-1]

	return
}
