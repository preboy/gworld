package lobby

import "sort"

type CardsType int

const (
	CardsTypeNIL          CardsType = iota + 0 // 非法
	CardsTypeJJ                                // 王炸
	CardsTypeA                                 // 单牌
	CardsTypeA_SEQ                             // 单牌(顺子)(5)
	CardsTypeAA_SEQ                            // 对牌(顺子)(3)
	CardsTypeAAA                               // 三不带
	CardsTypeAAA_SEQ                           // 三不带(顺子)
	CardsTypeAAAX                              // 三带单
	CardsTypeAAAX_SEQ                          // 三带单(顺子)
	CardsTypeAAAXX                             // 三带对
	CardsTypeAAAXX_SEQ                         // 三带对(顺子)
	CardsTypeAAAA                              // 炸弹
	CardsTypeAAAA_SEQ                          // 炸弹(顺子)
	CardsTypeAAAAXY                            // 四带二
	CardsTypeAAAAXY_SEQ                        // 四带二(顺子)
	CardsTypeAAAAXXYY                          // 四带对
	CardsTypeAAAAXXYY_SEQ                      // 四带对(顺子)
)

type Analyse struct {
	Cards map[int32][]Card // point -> cards
}

type CardsInfo struct {
	Type CardsType // 牌类型
	Max  int32     // 主牌最大点
	Len  int32     // 主牌长度(SEQ)
}

type AnalyseItem struct {
	Point int32
	Count int32
}

// ----------------------------------------------------------------------------

func NewAnalyse(cards []Card) *Analyse {
	a := &Analyse{
		Cards: map[int32][]Card{},
	}

	for _, c := range cards {
		p := c.Point()
		a.Cards[p] = append(a.Cards[p], c)
	}

	return a
}

func (self *Analyse) Points() (ret []*AnalyseItem) {
	for p, v := range self.Cards {
		ret = append(ret, &AnalyseItem{p, int32(len(v))})
	}

	// sort
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].Count != ret[j].Count {
			return ret[i].Count > ret[j].Count
		}

		return ret[i].Point > ret[j].Point
	})

	return
}

// ----------------------------------------------------------------------------

func cards_from_int32(cards []int32) (ret []Card, valid bool) {
	for _, v := range cards {
		c := NewCardFromValue(v)
		ret = append(ret, c)

		if !c.Valid() {
			return
		}
	}

	valid = true
	return
}

func cards_to_int32(cards []Card) (ret []int32) {
	for _, c := range cards {
		ret = append(ret, c.Value())
	}

	return
}

func cards_sort(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
}

func get_cards_info(cards []Card) *CardsInfo {
	cards_sort(cards)

	// a := NewAnalyse(cards)

	ci := &CardsInfo{
		Type: CardsTypeNIL,
	}

	switch len(cards) {

	case 1: // CardsTypeA
		ci.Max = cards[0].Point()

	case 2: // CardsTypeJJ
		if cards[0].Value() == CardValue_J2 && cards[1].Value() == CardValue_J1 {
			ci.Type = CardsTypeJJ
		}

	case 3: // CardsTypeAAA

	case 4: // CardsTypeAAAA CardsTypeAAAX

	case 5: // CardsTypeA_SEQ CardsTypeAAAXX

	case 6: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAAXY

	case 7: // CardsTypeA_SEQ

	case 8: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ

	case 9: // CardsTypeA_SEQ CardsTypeAAA_SEQ

	case 10: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAAXX_SEQ

	case 11: // CardsTypeA_SEQ

	case 12: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAAXY_SEQ

	case 14: // CardsTypeAA_SEQ

	case 15: // CardsTypeAAA_SEQ CardsTypeAAAXX_SEQ

	case 16: // CardsTypeAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ CardsTypeAAAAXXYY_SEQ

	case 18: // CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAAXY_SEQ

	case 20: // CardsTypeAA_SEQ CardsTypeAAAA_SEQ CardsTypeAAAXX_SEQ

	default:
		break
	}

	return ci
}
