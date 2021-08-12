package netmgr

import (
	"math/rand"

	"gworld/core"
	"gworld/core/log"
	"gworld/ddz/lobby/poker"
)

const (
	STRATEGY = 3 // AI策略 (1: 分析		2: 能大就大		3: 能不出就不出)
)

// ----------------------------------------------------------------------------
// ai

func (self *AILogic) ai_call() int32 {
	if STRATEGY == 1 {
		return rand.Int31n(4)
	}

	return rand.Int31n(4)
}

func (self *AILogic) ai_play(first bool) (cards []poker.Card) {

	switch STRATEGY {
	case 1:
		{
			a1 := poker.NewAnalyse(self.cards)
			a2 := poker.NewAnalyse(self.cards_left)

			// TODO  计算其它方位可能有的牌

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

	case 2:
		{

			if first {
				// todo 随机出一手

			} else {
				a1 := poker.NewAnalyse(self.cards)

				cards_prev := self.prev_play()
				if cards_prev == nil {
					panic("cards_prev == nil")
				}

				// 上一首不为Nil的牌
				ci := poker.CardsAnalyse(cards_prev)

				// 找到大过他的牌
				ret := a1.Exceed(ci)
				for i, v := range ret {
					log.Info("the <%d> Exceed er: %v", i, poker.CardsToString(v))
				}

				if len(ret) > 0 {
					cards = ret[0]
				}
			}

			return
		}

	case 3:
		{
			if first {
				poker.CardsSort(self.cards)
				l := len(self.cards)
				cards = self.cards[l-1:]
			}

			return
		}
	}

	return
}

// ----------------------------------------------------------------------------
// ai analyse

type section struct {
	init bool
	p1   int32
	p2   int32
}

func (self *section) push(v int32) bool {
	if !self.init {
		self.init = true

		self.p1 = v
		self.p2 = v

		return true
	}

	if v != self.p2+1 {
		return false
	}

	self.p2 = v

	return true
}

func (self *section) length() int32 {
	return self.p2 - self.p1 + 1
}

func (self *section) conform(l int32) bool {
	return self.length() >= l
}

// ----------------------------------------------------------------------------

type divide_type int32

const (
	divide_type_A divide_type = iota + 0
	divide_type_AA
	divide_type_AAA
	divide_type_AAAA
	divide_type_ABCDE
	divide_type_AABBCC
	divide_type_AAABBB
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
	divide_type_string := []string{
		"divide_type_A    	",
		"divide_type_AA   	",
		"divide_type_AAA  	",
		"divide_type_AAAA 	",
		"divide_type_ABCDE	",
		"divide_type_AABBCC	",
		"divide_type_AAABBB	",
	}

	ret := divide_type_string[self.dtype] + ": "

	ret += "{ "
	for _, v := range self.items {
		ret += poker.CardsToString(v)
		ret += ", "
	}
	ret += " }"

	return ret
}

func (self *divide_t) add(cards []poker.Card) {
	self.items = append(self.items, cards)
}

func (self *divide_t) keys() (ret []int32) {
	for _, v := range self.items {
		if len(v) > 0 {
			ret = append(ret, v[0].Point())
		}
	}

	core.SortInt32s(ret)
	return
}

func (self *divide_t) empty() bool {
	if len(self.items) == 0 {
		return true
	}

	for _, v := range self.items {
		if len(v) != 0 {
			return false
		}
	}

	return true
}

func (self *divide_t) all_cards() (ret []poker.Card) {
	for _, v := range self.items {
		ret = append(ret, v...)
	}
	return
}

func (self *divide_t) take_sequence(l int32) (ret [][]poker.Card) {
	if len(self.items) < int(l) {
		return
	}

	if self.dtype != divide_type_AA && self.dtype != divide_type_AAA && self.dtype != divide_type_AAAA {
		panic("error dtype in take_sequence")
	}

	var arr []*section

	// find
	{
		sec := &section{}
		for _, v := range self.keys() {
			if !sec.push(v) {

				b := sec.conform(l)
				if b {
					arr = append(arr, sec)
				}

				sec = &section{}
				if !b {
					sec.push(v)
				}
			}
		}

		if sec.conform(l) {
			arr = append(arr, sec)
		}
	}

	// extract
	for _, v := range arr {
		cards := self.extract(v.p1, v.p2)

		if len(cards) > 0 {
			ret = append(ret, cards)
		}
	}

	return
}

func (self *divide_t) extract(p1, p2 int32) (cards []poker.Card) {
	var items [][]poker.Card

	for _, v := range self.items {
		if len(v) == 0 {
			continue
		}

		p := v[0].Point()

		if p >= p1 && p <= p2 {
			cards = append(cards, v...)
		} else {
			items = append(items, v)
		}
	}

	self.items = items

	return
}

func (self *divide_t) merge_sequence(cards []poker.Card) []poker.Card {

REPEAT:

	for {
		if self.merge_internal() {
			break
		}
	}

	for _, c := range cards {
		if self.join_card(c) {
			cards = poker.CardsRemoveOne(cards, c)
			goto REPEAT
		}
	}

	return cards
}

func (self *divide_t) merge_internal() bool {
	l := len(self.items)
	if l < 2 {
		return true
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if self.join(i, j) {
				return false
			}
		}
	}

	return true
}

func (self *divide_t) join(i, j int) bool {
	s1 := self.items[i]
	s2 := self.items[j]

	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	s1_1 := s1[0]
	s1_2 := s1[len(s1)-1]
	s2_1 := s2[0]
	s2_2 := s2[len(s2)-1]

	var join bool
	var item []poker.Card

	if s1_2.Point()+1 == s2_1.Point() {
		item = append(s1, s2...)
		join = true
	}

	if s2_2.Point()+1 == s1_1.Point() {
		item = append(s2, s1...)
		join = true
	}

	if join {
		var items [][]poker.Card

		for k, v := range self.items {
			if k != i && k != j {
				items = append(items, v)
			}
		}

		items = append(items, item)
		self.items = items
	}

	return join
}

func (self *divide_t) join_card(c poker.Card) bool {
	for k, v := range self.items {
		l := len(v)
		if l == 0 {
			continue
		}

		if c.Point() == v[0].Point()-1 {
			self.items[k] = append([]poker.Card{c}, v...)
			return true
		}

		if c.Point() == v[l-1].Point()+1 {
			self.items[k] = append(v, c)
			return true
		}

	}

	return false
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
		self.divides[dtype] = &divide_t{
			dtype: dtype,
		}
	}

	return self.divides[dtype]
}

func (self *class_t) add(dtype divide_type, cards []poker.Card) {
	self.get(dtype).add(cards)
}

func (self *class_t) del(dtype divide_type) {
	delete(self.divides, dtype)
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
	ptr, ok := self.divides[divide_type_AA]
	if !ok {
		return
	}

	// 从现有的牌中抽出2顺
	for _, cards := range ptr.take_sequence(3) {
		self.add(divide_type_AABBCC, cards)
	}

	if ptr.empty() {
		self.del(ptr.dtype)
	}

	// TODO 甚至向其它地方(3带)借一对，能还连起2顺
}

func (self *class_t) merge_aaa() {
	// 能否连成飞机
	ptr := self.divides[divide_type_AAA]
	if ptr == nil {
		return
	}

	// 从现有的牌中抽出3顺
	for _, cards := range ptr.take_sequence(2) {
		self.add(divide_type_AAABBB, cards)
	}
}

func (self *class_t) merge_aaaa() {
	// DO NOTHING NOW
}

func (self *class_t) merge_abcde() {
	ptr := self.divides[divide_type_ABCDE]
	if ptr == nil {
		return
	}

	var cards []poker.Card
	p := self.divides[divide_type_A]
	if p != nil {
		cards = p.all_cards()
	}

	var items [][]poker.Card
	for _, c := range ptr.merge_sequence(cards) {
		items = append(items, []poker.Card{c})
	}

	if p != nil {
		p.items = items
	}
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

	c.merge_aa()
	c.merge_aaa()
	c.merge_aaaa()
	c.merge_abcde()

	return c
}

func cards_divide_aaa(cards []poker.Card) *class_t {
	c := new_class()

	cards = c.pull_aaa(cards)
	cards = c.pull_abcde(cards)
	cards = c.pull_aaaa(cards)
	cards = c.pull_aa(cards)
	cards = c.pull_a(cards)

	if len(cards) != 0 {
		panic("not empty")
	}

	c.merge_aa()
	c.merge_aaa()
	c.merge_aaaa()
	c.merge_abcde()

	return c
}
