package lobby

import (
	"fmt"
	"sort"
)

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

type AnalysedPoint struct {
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

func (self *Analyse) GetPoints() (ret []*AnalysedPoint) {
	for p, v := range self.Cards {
		ret = append(ret, &AnalysedPoint{p, int32(len(v))})
	}

	// sort
	if len(ret) > 1 {
		sort.Slice(ret, func(i, j int) bool {
			if ret[i].Count != ret[j].Count {
				return ret[i].Count > ret[j].Count
			}

			return ret[i].Point > ret[j].Point
		})
	}

	return
}

func (self *CardsInfo) ToString() string {
	return fmt.Sprintf("%v %v %v", self.Type, self.Max, self.Len)
}

// ----------------------------------------------------------------------------

func cards_from_int32(cards []int32) (ret []Card, valid bool) {
	for _, v := range cards {
		c := NewCardFromValue(v)
		if !c.Valid() {
			return
		}

		ret = append(ret, c)
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

	a := NewAnalyse(cards)
	points := a.GetPoints()
	cnt_points := int32(len(points))

	ci := &CardsInfo{
		Type: CardsTypeNIL,
	}

	cnt_cards := int32(len(cards))

	switch cnt_cards {

	case 1: // CardsTypeA
		ci.Type = CardsTypeA
		ci.Max = cards[0].Point()

	case 2: // CardsTypeJJ
		if cards[0].Value() == CardValue_J2 && cards[1].Value() == CardValue_J1 {
			ci.Type = CardsTypeJJ
		}

	case 3: // CardsTypeAAA
		if cnt_points == 1 {
			ci.Type = CardsTypeAAA
			ci.Max = points[0].Point
		}

	case 4: // CardsTypeAAAA CardsTypeAAAX
		if cnt_points == 1 {
			ci.Type = CardsTypeAAAA
			ci.Max = points[0].Point
			break
		}

		if cnt_points == 2 && points[0].Count == 3 {
			ci.Type = CardsTypeAAAX
			ci.Max = points[0].Point
		}

	case 5: // CardsTypeA_SEQ CardsTypeAAAXX
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if cnt_points == 2 && points[0].Count == 3 {
			ci.Type = CardsTypeAAAXX
			ci.Max = points[0].Point
		}

	case 6: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAAXY
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if (cnt_points == 2 || cnt_points == 3) && points[0].Count == 4 {
			ci.Type = CardsTypeAAAAXY
			ci.Max = points[0].Point
			break
		}

	case 7: // CardsTypeA_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

	case 8: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ CardsTypeAAAAXXYY_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAX_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAAXXYY_SEQ(points, cnt_cards, ci) {
			break
		}

	case 9: // CardsTypeA_SEQ CardsTypeAAA_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAA_SEQ(points, cnt_cards, ci) {
			break
		}

	case 10: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAAXX_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAXX_SEQ(points, cnt_cards, ci) {
			break
		}

	case 11: // CardsTypeA_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

	case 12: // CardsTypeA_SEQ CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ CardsTypeAAAAXY_SEQ
		if Is_CardsTypeA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAX_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAAXY_SEQ(points, cnt_cards, ci) {
			break
		}

	case 14: // CardsTypeAA_SEQ
		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

	case 15: // CardsTypeAAA_SEQ CardsTypeAAAXX_SEQ
		if Is_CardsTypeAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAXX_SEQ(points, cnt_cards, ci) {
			break
		}

	case 16: // CardsTypeAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ CardsTypeAAAAXXYY_SEQ
		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAX_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAAXXYY_SEQ(points, cnt_cards, ci) {
			break
		}

	case 18: // CardsTypeAA_SEQ CardsTypeAAA_SEQ CardsTypeAAAAXY_SEQ
		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAAXY_SEQ(points, cnt_cards, ci) {
			break
		}

	case 20: // CardsTypeAA_SEQ CardsTypeAAAX_SEQ CardsTypeAAAA_SEQ CardsTypeAAAXX_SEQ
		if Is_CardsTypeAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAX_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAA_SEQ(points, cnt_cards, ci) {
			break
		}

		if Is_CardsTypeAAAXX_SEQ(points, cnt_cards, ci) {
			break
		}

	default:
		break
	}

	return ci
}

func Is_CardsTypeA_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards

	b := cnt_points == seq_length &&
		points[0].Point-points[cnt_points-1].Point == seq_length-1

	if b {
		ci.Type = CardsTypeA_SEQ
		ci.Max = points[0].Point
		ci.Len = seq_length
	}

	return b
}

func Is_CardsTypeAA_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 2

	if cnt_points != seq_length {
		return false
	}

	for i := int32(0); i < cnt_points; i++ {
		if points[i].Count != 2 {
			return false
		}
	}

	b := points[0].Point-points[cnt_points-1].Point == seq_length-1

	if b {
		ci.Type = CardsTypeAA_SEQ
		ci.Max = points[0].Point
		ci.Len = seq_length
	}

	return b
}

func Is_CardsTypeAAA_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 3

	if cnt_points != seq_length {
		return false
	}

	for i := int32(0); i < cnt_points; i++ {
		if points[i].Count != 3 {
			return false
		}
	}

	b := points[0].Point-points[cnt_points-1].Point == seq_length-1

	if b {
		ci.Type = CardsTypeAAA_SEQ
		ci.Max = points[0].Point
		ci.Len = seq_length
	}

	return b
}

func Is_CardsTypeAAAX_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 4

	if cnt_points < seq_length {
		return false
	}

	for i := int32(0); i < seq_length; i++ {
		if points[i].Count < 3 {
			return false
		}
	}

	if points[0].Point-points[seq_length-1].Point != seq_length-1 {
		return false
	}

	ci.Type = CardsTypeAAAX_SEQ
	ci.Max = points[0].Point
	ci.Len = seq_length

	// example: 3333 5555 6666
	return true
}

func Is_CardsTypeAAAA_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 4

	if cnt_points != seq_length {
		return false
	}

	for i := int32(0); i < seq_length; i++ {
		if points[i].Count < 4 {
			return false
		}
	}

	if points[0].Point-points[seq_length-1].Point != seq_length-1 {
		return false
	}

	ci.Type = CardsTypeAAAA_SEQ
	ci.Max = points[0].Point
	ci.Len = seq_length

	return true
}

func Is_CardsTypeAAAXX_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 5

	if cnt_points <= seq_length {
		return false
	}

	for i := int32(0); i < seq_length; i++ {
		if points[i].Count < 3 {
			return false
		}
	}

	for i := int32(seq_length); i < seq_length*2; i++ {
		if points[i].Count < 2 {
			return false
		}
	}

	if points[0].Point-points[seq_length-1].Point != seq_length-1 {
		return false
	}

	// 不考虑副牌是炸弹的情况

	ci.Type = CardsTypeAAAXX_SEQ
	ci.Max = points[0].Point
	ci.Len = seq_length

	return true
}

func Is_CardsTypeAAAAXY_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 6

	if cnt_points <= seq_length {
		return false
	}

	for i := int32(0); i < seq_length; i++ {
		if points[i].Count < 4 {
			return false
		}
	}

	if points[0].Point-points[seq_length-1].Point != seq_length-1 {
		return false
	}

	// 不考虑副牌凑成炸弹的情况

	ci.Type = CardsTypeAAAAXY_SEQ
	ci.Max = points[0].Point
	ci.Len = seq_length

	return true
}

func Is_CardsTypeAAAAXXYY_SEQ(points []*AnalysedPoint, cnt_cards int32, ci *CardsInfo) bool {
	cnt_points := int32(len(points))
	seq_length := cnt_cards / 8

	if cnt_points <= seq_length {
		return false
	}

	for i := int32(0); i < seq_length; i++ {
		if points[i].Count < 4 {
			return false
		}
	}

	for i := int32(seq_length); i < seq_length*3; i++ {
		if points[i].Count < 2 {
			return false
		}
	}

	if points[0].Point-points[seq_length-1].Point != seq_length-1 {
		return false
	}

	// 不考虑副牌凑成炸弹的情况

	ci.Type = CardsTypeAAAAXXYY_SEQ
	ci.Max = points[0].Point
	ci.Len = seq_length

	return true
}
