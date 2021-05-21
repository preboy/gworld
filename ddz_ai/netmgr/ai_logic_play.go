package netmgr

import (
	"gworld/core/log"
	"gworld/ddz/lobby/poker"
)

type data struct {
	left_count int32        // 剩余牌数量
	left_cards []poker.Card // 自己(空)  其它人（可能的牌）
	deal_cards []poker.Card // 出过的牌
}

type hand struct {
	pos   int32
	cards []poker.Card
}

type round struct {
	hands []*hand
}

// ----------------------------------------------------------------------------
// event

func (self *AILogic) on_init() {
	self.plrs = map[int32]*data{}

	for i := int32(0); i < 3; i++ {
		self.plrs[i] = &data{
			left_count: 17,
		}
	}

	self.left_cards, _ = poker.CardsRemove(poker.NewPoker(), self.cards)
}

func (self *AILogic) on_calc() {
	self.plrs[self.lord_pos].left_count += 3

	if self.lord_pos == self.pos {
		self.left_cards, _ = poker.CardsRemove(self.left_cards, self.lord_cards)
	}
}

func (self *AILogic) on_play(pos int32, first bool, cards []poker.Card) {
	if first {
		self.rounds = append(self.rounds, &round{})
	}

	self.add_play(pos, cards)

	if self.pos == pos {
		self.cards, _ = poker.CardsRemove(self.cards, cards)
		return
	}

	if len(cards) != 0 {
		self.left_cards, _ = poker.CardsRemove(self.left_cards, cards)
		self.plrs[pos].deal_cards = append(self.plrs[pos].deal_cards, cards...)
		self.plrs[pos].left_count -= int32(len(cards))
	}
}

func (self *AILogic) play(first bool) (cards []poker.Card, ok bool) {

	a1 := poker.NewAnalyse(self.cards)
	a2 := poker.NewAnalyse(self.left_cards)

	_ = a2

	if first {
		// 首出

	} else {

		// 跟牌

		cards_prev := self.prev_play()
		if cards_prev == nil {
			panic("cards_prev == nil")
		}

		// 上一首不为Nil的牌
		ci := poker.CardsAnalyse(cards_prev)

		// 找到大过他的牌
		ret := a1.Exceed(ci)
		for i, v := range ret {
			log.Info("the <%d> groups: %v", i, poker.CardsToString(v))
		}
	}

	return
}

// ----------------------------------------------------------------------------
// local

func (self *AILogic) is_lord() bool {
	return self.pos == self.lord_pos
}

func (self *AILogic) get_friend_pos() int32 {
	if ai.is_lord() {
		panic("is lord")
	}

	for i := int32(0); i < 3; i++ {
		if i == self.lord_pos || i == self.pos {
			continue
		}

		return i
	}

	panic("not found friend")
}

func (self *AILogic) add_play(pos int32, cards []poker.Card) {
	l := len(self.rounds)
	r := self.rounds[l-1]
	r.hands = append(r.hands, &hand{pos, cards})
}

func (self *AILogic) prev_play() []poker.Card {
	l := len(self.rounds)
	r := self.rounds[l-1]

	for i := len(r.hands) - 1; i >= 0; i-- {
		if len(r.hands[i].cards) != 0 {
			return r.hands[i].cards
		}
	}

	return nil
}

// ----------------------------------------------------------------------------

type divide_type int32

const (
	divide_type_A divide_type = iota + 0
	divide_type_AA
	divide_type_AAA
	divide_type_AAAA
	divide_type_ABCDE
)

var (
	divide_string = []string{
		"divide_type_A    ",
		"divide_type_AA   ",
		"divide_type_AAA  ",
		"divide_type_AAAA ",
		"divide_type_ABCDE",
	}
)

// ----------------------------------------------------------------------------

type divide_t struct {
	dtype divide_type
	items [][]poker.Card
}

type class_t struct {
	divides map[divide_type]*divide_t
}

// ----------------------------------------------------------------------------
// divide_t
func (self *divide_t) dump() string {
	ret := divide_string[self.dtype] + ": "

	ret += "{ "
	for _, v := range self.items {
		ret += poker.CardsToString(v)
		ret += ", "
	}
	ret += " }"

	return ret
}

// ----------------------------------------------------------------------------
// class_t

func new_class() *class_t {
	return &class_t{
		divides: map[divide_type]*divide_t{},
	}
}

func (self *class_t) get(dtype divide_type) *divide_t {
	if self.divides[dtype] == nil {
		self.divides[dtype] = &divide_t{dtype: dtype}
	}

	return self.divides[dtype]
}

func (self *class_t) add(dtype divide_type, cards []poker.Card) {
	d := self.get(dtype)
	d.items = append(d.items, cards)
}

func (self *class_t) pull_a(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 1 {
			self.add(divide_type_A, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (self *class_t) pull_aa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 2 {
			self.add(divide_type_AA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (self *class_t) pull_aaa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 3 {
			self.add(divide_type_AAA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (self *class_t) pull_aaaa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 4 {
			self.add(divide_type_AAAA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (self *class_t) pull_abcde(cards []poker.Card) []poker.Card {
	for i := poker.CardPoint_7; i <= poker.CardPoint_A; i++ {
		a := poker.NewAnalyse(cards)
		seq := a.GetSeq(i, 5, 1)
		if len(seq) != 0 {
			self.add(divide_type_ABCDE, seq)
			cards, _ = poker.CardsRemove(cards, seq)
		}
	}

	return cards
}

func (self *class_t) merge_aa() {
	// 现有的牌能否连起2顺
	// 甚至向其它地方(3带)借一对，能还连起2顺
}

func (self *class_t) merge_aaa() {
	// 能否连成飞机
}

func (self *class_t) merge_aaaa() {
	// DO NOTHING NOW
}

func (self *class_t) merge_abcde() {
	// 1 can ?

	// 2 need ?
}

func (self *class_t) merge() {
	self.merge_aa()
	self.merge_aaa()
	self.merge_aaaa()
	self.merge_abcde()
}

func (self *class_t) evaluate() int32 {
	// 评分：
	// 单牌、单对越少越好

	return 0
}

func (self *class_t) dump() (ret string) {
	ret += "\n"

	for _, v := range self.divides {
		ret += "\t"
		ret += v.dump()
		ret += "\n"
	}

	return
}

// ----------------------------------------------------------------------------

func cards_divide(cards []poker.Card) (ret []*class_t) {
	// 1
	c := cards_divide_abdef(cards)
	if c != nil {
		ret = append(ret, c)
	}

	// 2
	c = cards_divide_aaa(cards)
	if c != nil {
		ret = append(ret, c)
	}

	return
}

func cards_divide_abdef(cards []poker.Card) *class_t {
	c := new_class()

	cards = c.pull_abcde(cards)
	cards = c.pull_aaaa(cards)
	cards = c.pull_aaa(cards)
	cards = c.pull_aa(cards)
	cards = c.pull_a(cards)

	if len(cards) != 0 {
		panic("not empty")
	}

	c.merge()

	return c
}

func cards_divide_aaa(cards []poker.Card) *class_t {

	return nil
}
