package poker

import "sort"

// ----------------------------------------------------------------------------
// Analyse for ai

// 找到大过ci的牌
func (self *Analyse) Exceed(ci *CardsInfo) (ret [][]Card) {
	if ci.Type == CardsTypeNIL {
		panic("CardsTypeNIL")
	}

	if ci.Type == CardsTypeJJ {
		return
	}

	switch ci.Type {
	case CardsTypeA:
		for p, v := range self.Cards {
			if p > ci.Max && len(v) > 0 {
				var cards []Card = []Card{v[0]}
				ret = append(ret, cards)
			}
		}

	case CardsTypeA_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 1)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAA_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 2)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAA:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 3)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAA_SEQ:
		// todo

	default:
		break
	}

	return
}

// ----------------------------------------------------------------------------
// aux

func (self *Analyse) get_seq(point int32, length int32, count int) (ret []Card) {
	for i := point - length + 1; i <= point; i++ {
		if len(self.Cards[i]) < count {
			return nil
		} else {
			ret = append(ret, self.Cards[i][:count]...)
		}
	}

	return
}

func (self *Analyse) get_points_sorted() (ret []int32) {
	for p, v := range self.Cards {
		if len(v) > 0 {
			ret = append(ret, p)
		}
	}

	if len(ret) == 0 {
		return
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})

	return
}
