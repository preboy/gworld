package poker

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
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 1)
			if len(cards) > 0 {
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
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 3)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAX:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAX_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAXX:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAXX_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAA:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 4)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAA_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 4)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXY:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A(), NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXY_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A(), NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXXYY:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, 1, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA(), NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXXYY_SEQ:
		for i := ci.Max + 1; i <= CardPoint_A; i++ {
			cards := self.get_seq(i, ci.Len, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA(), NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	default:
		break
	}

	// 其它的牌型把炸弹加上
	if !ci.IsBomb() {
		for _, v := range self.get_bombs() {
			ret = append(ret, v)
		}
	} else {
		jj := self.get_jj()
		if len(jj) > 0 {
			ret = append(ret, jj)
		}
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

func (self *Analyse) get_jj() (ret []Card) {
	if len(self.Cards[CardPoint_J1]) == 1 && len(self.Cards[CardPoint_J2]) == 1 {
		ret = append(ret, self.Cards[CardPoint_J1][0], self.Cards[CardPoint_J2][0])
	}

	return
}

func (self *Analyse) get_bombs() (ret [][]Card) {
	for _, v := range self.Cards {
		if len(v) == 4 {
			ret = append(ret, v)
		}
	}

	jj := self.get_jj()
	if len(jj) > 0 {
		ret = append(ret, jj)
	}

	return
}
