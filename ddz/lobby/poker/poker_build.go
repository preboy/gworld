package poker

import (
	"math/rand"
)

const (
	cards_max_number = 17
)

type (
	card_wall map[int32][]Card
)

// ----------------------------------------------------------------------------
// poker_build

func (self card_wall) init() {
	// 3 ~ A * 4
	for p := CardPoint_3; p <= CardPoint_A; p++ {
		var cards []Card
		for c := CardColor_Heart; c <= CardColor_Club; c++ {
			cards = append(cards, NewCard(c, p))
		}
		cards_random(cards)
		self[p] = cards
	}

	// 2 * 4
	{
		var cards []Card
		for c := CardColor_Heart; c <= CardColor_Club; c++ {
			cards = append(cards, NewCard(c, CardPoint_2))
		}
		cards_random(cards)
		self[CardPoint_2] = cards
	}

	// joker
	{
		c := NewCardFromValue(CardValue_J1)
		p := c.Point()
		self[p] = append(self[p], c)
	}

	// joker
	{
		c := NewCardFromValue(CardValue_J2)
		p := c.Point()
		self[p] = append(self[p], c)
	}
}

func (self card_wall) exist_seq(p2 int32, length int32, cnt int) bool {
	for i := p2 - length + 1; i <= p2; i++ {
		if len(self[i]) < cnt {
			return false
		}
	}
	return true
}

func (self card_wall) expect_abcde(length int32) (ret []Card) {
	if length < 5 {
		return
	}

	p1 := 3 + length - 1
	p2 := int32(14)

	set := []int32{}
	for i := p1; i <= p2; i++ {
		if self.exist_seq(i, length, 1) {
			set = append(set, i)
		}
	}

	if len(set) == 0 {
		return self.expect_abcde(length - 1)
	}

	r := rand.Intn(len(set))
	v := set[r]

	for i := v - length + 1; i <= v; i++ {
		l, c := cards_remove(self[i])
		ret = append(ret, c)
		self[i] = l
	}

	return
}

func (self card_wall) expect_aabbcc_seq(length int32) (ret []Card) {
	if length < 3 {
		return
	}

	p1 := 3 + length - 1
	p2 := int32(14)

	set := []int32{}
	for i := p1; i <= p2; i++ {
		if self.exist_seq(i, length, 2) {
			set = append(set, i)
		}
	}

	if len(set) == 0 {
		return self.expect_aabbcc_seq(length - 1)
	}

	r := rand.Intn(len(set))
	v := set[r]

	for i := v - length + 1; i <= v; i++ {
		for j := 0; j < 2; j++ {
			l, c := cards_remove(self[i])
			ret = append(ret, c)
			self[i] = l
		}
	}

	return
}

func (self card_wall) expect_aaa_seq(length int32) (ret []Card) {
	if length < 1 {
		return
	}

	p1 := 3 + length - 1
	p2 := int32(14)

	set := []int32{}
	for i := p1; i <= p2; i++ {
		if self.exist_seq(i, length, 3) {
			set = append(set, i)
		}
	}

	if len(set) == 0 {
		return self.expect_aaa_seq(length - 1)
	}

	r := rand.Intn(len(set))
	v := set[r]

	for i := v - length + 1; i <= v; i++ {
		for j := 0; j < 3; j++ {
			l, c := cards_remove(self[i])
			ret = append(ret, c)
			self[i] = l
		}
	}

	return
}

func (self card_wall) expect_aaaa() (ret []Card) {

	set := []int32{}

	for i := CardPoint_3; i <= CardPoint_2; i++ {
		if i == 15 {
			continue
		}

		if self.exist_seq(i, 1, 4) {
			set = append(set, i)
		}
	}

	if len(set) == 0 {
		return
	}

	r := rand.Intn(len(set))
	v := set[r]

	for i := 0; i < 4; i++ {
		l, c := cards_remove(self[v])
		ret = append(ret, c)
		self[v] = l
	}

	return
}

// ----------------------------------------------------------------------------
// poker_build

type poker_build struct {
	wall card_wall
	draw [][]Card
}

func NewPokerBuild() *poker_build {
	pb := &poker_build{
		wall: make(card_wall),
	}
	pb.wall.init()

	return pb
}

func cards_append(src []Card, add []Card) (r []Card, b bool) {
	if len(src)+len(add) > cards_max_number {
		return src, false
	}

	r = append(src, add...)
	b = true

	return
}

func (self *poker_build) Finish() (ret []Card) {

	var (
		seats      = [3][]Card{}
		cards_left = []Card{}
	)

	// draw -> seat
	for _, cards := range self.draw {
		seq := []int{0, 1, 2}
		int_random(seq)

		b := false
		for i := 0; i < len(seq); i++ {
			p := seq[i]

			r := []Card{}
			r, b = cards_append(seats[p], cards)
			if b {
				seats[p] = r
				break
			}
		}

		if !b {
			cards_left = append(cards_left, cards...)
		}
	}

	// wall -> seat
	for _, arr := range self.wall {
		if len(arr) > 0 {
			cards_left = append(cards_left, arr...)
		}
	}

	cards_random(cards_left)

	for i := 0; i < 3; i++ {

		l1 := 17 - len(seats[i])
		l2 := len(cards_left)

		if l2 == 0 {
			break
		}

		seats[i] = append(seats[i], cards_left[l2-l1:l2]...)
		cards_left = cards_left[:l2-l1]
	}

	for i := 0; i < 3; i++ {
		ret = append(ret, seats[i]...)
	}

	ret = append(ret, cards_left...)
	return
}

func (self *poker_build) Build_ABCDE(count int32, length int32) bool {
	// args checking
	{
		if count <= 0 || count > 5 {
			return false
		}

		if length < 5 || length > 12 {
			return false
		}
	}

	for i := int32(0); i < count; i++ {
		l := random_angle_45(5, length)

		cards := self.wall.expect_abcde(l)
		if len(cards) != 0 {
			self.draw = append(self.draw, cards)
		}
	}

	return true
}

func (self *poker_build) Build_AABBCC(count int32, length int32) bool {
	// args checking
	{
		if count <= 0 || count > 4 {
			return false
		}

		if length < 3 || length > 8 {
			return false
		}
	}

	for i := int32(0); i < count; i++ {
		l := random_angle_45(3, length)

		cards := self.wall.expect_aabbcc_seq(l)
		if len(cards) != 0 {
			self.draw = append(self.draw, cards)
		}
	}

	return true
}

func (self *poker_build) Build_AAABBB(count int32, length int32) bool {
	// args checking
	{
		if count <= 0 || count > 4 {
			return false
		}

		if length < 1 || length > 5 {
			return false
		}
	}

	for i := int32(0); i < count; i++ {
		l := random_angle_45(1, length)

		cards := self.wall.expect_aaa_seq(l)
		if len(cards) != 0 {
			self.draw = append(self.draw, cards)
		}
	}

	return true
}

// 不包括双王
func (self *poker_build) Build_AAAA(count int32) bool {

	if count <= 0 || count > 12 {
		return false
	}

	for i := int32(0); i < count; i++ {
		cards := self.wall.expect_aaaa()
		if len(cards) != 0 {
			self.draw = append(self.draw, cards)
		}
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

func cards_remove(cards []Card) ([]Card, Card) {
	l := len(cards)
	c := cards[l-1]

	left := cards[:l-1]

	return left, c
}
