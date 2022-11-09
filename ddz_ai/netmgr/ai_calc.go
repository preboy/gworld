package netmgr

import (
	"math/rand"

	"gworld/core"
	"gworld/core/log"
	"gworld/ddz/lobby/poker"
)

type STRATEGY int

const (
	STRATEGY_CALC STRATEGY = 1 + iota // 分析
	STRATEGY_MUST                     // 能大就大
	STRATEGY_PASS                     // 能不出就不出
)

const (
	STRATEGY_METHOD = STRATEGY_MUST
)

// ----------------------------------------------------------------------------
// ai

func (ai *AILogic) ai_call() int32 {
	if STRATEGY_METHOD == STRATEGY_CALC {
		return rand.Int31n(4)
	}

	return rand.Int31n(4)
}

func (ai *AILogic) ai_play(first bool) (cards []poker.Card) {

	switch STRATEGY_METHOD {

	case STRATEGY_CALC:
		{
			a1 := poker.NewAnalyse(ai.cards)
			a2 := poker.NewAnalyse(ai.cards_left)

			// TODO  计算其它方位可能有的牌

			_ = a2

			if first {
				// 首出

			} else {

				// 跟牌

				cards_prev := ai.prev_play()
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
		}

	case STRATEGY_MUST:
		{
			a1 := poker.NewAnalyse(ai.cards)

			if first {
				// TODO 随机出一手，尽量能回收，切不把手里的牌搞乱
				c := cards_divide_abdef(ai.cards)
				cards = c.first()
			} else {

				cards_prev := ai.prev_play()
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

				// TODO 处理带牌的情况

				if len(ret) > 0 {
					cards = ret[0]
				}
			}
		}

	case STRATEGY_PASS:
		{
			if first {
				poker.CardsSort(ai.cards)
				l := len(ai.cards)
				cards = ai.cards[l-1:]
			}
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

func (ai *section) push(v int32) bool {
	if !ai.init {
		ai.init = true

		ai.p1 = v
		ai.p2 = v

		return true
	}

	if v != ai.p2+1 {
		return false
	}

	ai.p2 = v

	return true
}

func (ai *section) length() int32 {
	return ai.p2 - ai.p1 + 1
}

func (ai *section) conform(l int32) bool {
	return ai.length() >= l
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
func (ai *divide_t) dump() string {
	divide_type_string := []string{
		"divide_type_A    	",
		"divide_type_AA   	",
		"divide_type_AAA  	",
		"divide_type_AAAA 	",
		"divide_type_ABCDE	",
		"divide_type_AABBCC	",
		"divide_type_AAABBB	",
	}

	ret := divide_type_string[ai.dtype] + ": "

	ret += "{ "
	for _, v := range ai.items {
		ret += poker.CardsToString(v)
		ret += ", "
	}
	ret += " }"

	return ret
}

func (ai *divide_t) add(cards []poker.Card) {
	ai.items = append(ai.items, cards)
}

func (ai *divide_t) keys() (ret []int32) {
	for _, v := range ai.items {
		if len(v) > 0 {
			ret = append(ret, v[0].Point())
		}
	}

	core.SortInt32s(ret)
	return
}

func (ai *divide_t) empty() bool {
	if len(ai.items) == 0 {
		return true
	}

	for _, v := range ai.items {
		if len(v) != 0 {
			return false
		}
	}

	return true
}

func (ai *divide_t) all_cards() (ret []poker.Card) {
	for _, v := range ai.items {
		ret = append(ret, v...)
	}
	return
}

func (ai *divide_t) take_sequence(l int32) (ret [][]poker.Card) {
	if len(ai.items) < int(l) {
		return
	}

	if ai.dtype != divide_type_AA && ai.dtype != divide_type_AAA && ai.dtype != divide_type_AAAA {
		panic("error dtype in take_sequence")
	}

	var arr []*section

	// find
	{
		sec := &section{}
		for _, v := range ai.keys() {
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
		cards := ai.extract(v.p1, v.p2)

		if len(cards) > 0 {
			ret = append(ret, cards)
		}
	}

	return
}

func (ai *divide_t) extract(p1, p2 int32) (cards []poker.Card) {
	var items [][]poker.Card

	for _, v := range ai.items {
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

	ai.items = items

	return
}

func (ai *divide_t) merge_sequence(cards []poker.Card) []poker.Card {

REPEAT:

	for {
		if ai.merge_internal() {
			break
		}
	}

	for _, c := range cards {
		if ai.join_card(c) {
			cards = poker.CardsRemoveOne(cards, c)
			goto REPEAT
		}
	}

	return cards
}

func (ai *divide_t) merge_internal() bool {
	l := len(ai.items)
	if l < 2 {
		return true
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if ai.join(i, j) {
				return false
			}
		}
	}

	return true
}

func (ai *divide_t) join(i, j int) bool {
	s1 := ai.items[i]
	s2 := ai.items[j]

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

		for k, v := range ai.items {
			if k != i && k != j {
				items = append(items, v)
			}
		}

		items = append(items, item)
		ai.items = items
	}

	return join
}

func (ai *divide_t) join_card(c poker.Card) bool {
	for k, v := range ai.items {
		l := len(v)
		if l == 0 {
			continue
		}

		if c.Point() == v[0].Point()-1 {
			ai.items[k] = append([]poker.Card{c}, v...)
			return true
		}

		if c.Point() == v[l-1].Point()+1 {
			ai.items[k] = append(v, c)
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

func (ai *class_t) get(dtype divide_type) *divide_t {
	if ai.divides[dtype] == nil {
		ai.divides[dtype] = &divide_t{
			dtype: dtype,
		}
	}

	return ai.divides[dtype]
}

func (ai *class_t) add(dtype divide_type, cards []poker.Card) {
	ai.get(dtype).add(cards)
}

func (ai *class_t) del(dtype divide_type) {
	delete(ai.divides, dtype)
}

func (ai *class_t) pull_a(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 1 {
			ai.add(divide_type_A, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (ai *class_t) pull_aa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 2 {
			ai.add(divide_type_AA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (ai *class_t) pull_aaa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 3 {
			ai.add(divide_type_AAA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (ai *class_t) pull_aaaa(cards []poker.Card) []poker.Card {
	a := poker.NewAnalyse(cards)

	for _, set := range a.Cards {
		if len(set) == 4 {
			ai.add(divide_type_AAAA, set)
			cards, _ = poker.CardsRemove(cards, set)
		}
	}

	return cards
}

func (ai *class_t) pull_abcde(cards []poker.Card) []poker.Card {
	for i := poker.CardPoint_7; i <= poker.CardPoint_A; i++ {
		a := poker.NewAnalyse(cards)
		seq := a.GetSeq(i, 5, 1)
		if len(seq) != 0 {
			ai.add(divide_type_ABCDE, seq)
			cards, _ = poker.CardsRemove(cards, seq)
		}
	}

	return cards
}

func (ai *class_t) merge_aabbcc() {
	ptr := ai.divides[divide_type_AA]
	if ptr == nil {
		return
	}

	// 从现有的牌中抽出2顺
	for _, cards := range ptr.take_sequence(3) {
		ai.add(divide_type_AABBCC, cards)
	}

	if ptr.empty() {
		ai.del(ptr.dtype)
	}

	// TODO 甚至向其它地方(3带)借一对，能还连起2顺
}

func (ai *class_t) merge_aaabbb() {
	// 能否连成飞机
	ptr := ai.divides[divide_type_AAA]
	if ptr == nil {
		return
	}

	// 从现有的牌中抽出3顺
	for _, cards := range ptr.take_sequence(2) {
		ai.add(divide_type_AAABBB, cards)
	}
}

func (ai *class_t) merge_aaaa() {
	// DO NOTHING NOW
}

func (ai *class_t) merge_abcde() {
	ptr := ai.divides[divide_type_ABCDE]
	if ptr == nil {
		return
	}

	var cards []poker.Card
	p := ai.divides[divide_type_A]
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

func (ai *class_t) evaluate() int {

	// 评分：
	// 单牌、单对越少越好
	// NOTE 暂时未考虑已出牌的情况

	single_card_count := 0

	{
		if ai.divides[divide_type_A] != nil {
			for _, v := range ai.divides[divide_type_A].items {
				if len(v) == 0 {
					continue
				}

				if int32(v[0]) == poker.CardValue_J1 {
					continue
				}

				single_card_count++
			}
		}

		if ai.divides[divide_type_AA] != nil {
			for _, v := range ai.divides[divide_type_AA].items {
				if len(v) != 0 {
					single_card_count++
				}

				// 对二， 对王 等最大的单牌不急算入内
				if int32(v[0]) == poker.CardPoint_2 {
					continue
				}

				single_card_count++
			}
		}

		if ai.divides[divide_type_AAA] != nil {
			for _, v := range ai.divides[divide_type_AAA].items {
				if len(v) != 0 {
					continue
				}

				single_card_count--

				if int32(v[0]) >= poker.CardPoint_J {
					continue
				}

				single_card_count++
			}
		}

		if ai.divides[divide_type_ABCDE] != nil {
			for _, v := range ai.divides[divide_type_ABCDE].items {
				if len(v) == 0 {
					continue
				}

				if len(v) > 6 {
					continue
				}

				if len(v) == 5 {
					if int32(v[0]) >= poker.CardPoint_7 {
						continue
					}
				}

				if len(v) == 6 {
					if int32(v[0]) >= poker.CardPoint_6 {
						continue
					}
				}

				single_card_count++
			}
		}
	}

	return single_card_count
}

func (ai *class_t) dump() (ret string) {
	ret += "\n"

	for _, v := range ai.divides {
		ret += "\t"
		ret += v.dump()
		ret += "\n"
	}

	return
}

// 首出:尽量找出能回手的牌
// 是否有更大的牌（或者自己就是最大的牌）
func (ai *class_t) first() []poker.Card {

	// 顺子

	// divide_type_ABCDE
	// divide_type_AABBCC
	// divide_type_AAABBB

	if ai.divides[divide_type_ABCDE] != nil {
		for _, cards := range ai.divides[divide_type_ABCDE].items {
			return cards
		}
	}

	// 3带1
	if ai.divides[divide_type_AABBCC] != nil {
		for _, cards := range ai.divides[divide_type_AABBCC].items {
			return cards
		}
	}

	if ai.divides[divide_type_AAABBB] != nil {
		for _, cards := range ai.divides[divide_type_AAABBB].items {
			return cards
		}
	}

	if ai.divides[divide_type_AAA] != nil {
		for _, cards := range ai.divides[divide_type_AAA].items {
			return cards
		}
	}

	if ai.divides[divide_type_AAA] != nil {
		for _, cards := range ai.divides[divide_type_AAA].items {
			return cards
		}
	}

	if ai.divides[divide_type_AAAA] != nil {
		for _, cards := range ai.divides[divide_type_AAAA].items {
			return cards
		}
	}

	if ai.divides[divide_type_A] != nil {
		for _, cards := range ai.divides[divide_type_A].items {
			return cards
		}
	}

	if ai.divides[divide_type_AA] != nil {
		for _, cards := range ai.divides[divide_type_AA].items {
			return cards[0:]
		}
	}

	// 实在没有就出一张单牌

	panic("未找到牌型: " + ai.dump())
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

	c.merge_aabbcc()
	c.merge_aaabbb()
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

	c.merge_aabbcc()
	c.merge_aaabbb()
	c.merge_aaaa()
	c.merge_abcde()

	return c
}
