package poker

// ----------------------------------------------------------------------------
// Analyse for ai

// 寻找大过对方的牌（粗暴）
func (self *Analyse) ExceedRough(ci *CardsInfo) (ret [][]Card) {
	if ci.Type == CardsTypeNIL {
		panic("CardsTypeNIL")
	}

	if ci.Type == CardsTypeJJ {
		return
	}

	switch ci.Type {

	case CardsTypeA:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 1)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeA_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 1)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAA_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 2)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAA:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 3)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAA_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 3)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAX:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAX_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAXX:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAXX_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 3)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAA:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 4)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAA_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 4)
			if len(cards) > 0 {
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXY:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A(), NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXY_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_A(), NewCard_PlaceHolder_A())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXXYY:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, 1, 4)
			if len(cards) > 0 {
				cards = append(cards, NewCard_PlaceHolder_AA(), NewCard_PlaceHolder_AA())
				ret = append(ret, cards)
			}
		}

	case CardsTypeAAAAXXYY_SEQ:
		for i := ci.MainMax + 1; i <= CardPoint_A; i++ {
			cards := self.GetSeq(i, ci.MainLen, 4)
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
		for _, v := range self.GetBombs() {
			ret = append(ret, v)
		}
	} else {
		jj := self.GetJJ()
		if len(jj) > 0 {
			ret = append(ret, jj)
		}
	}

	return
}

// 寻找大过对方的牌（优雅的）
func (self *Analyse) Exceed(ci *CardsInfo) (ret [][]Card) {
	return self.ExceedRough(ci)
}

// ----------------------------------------------------------------------------
// aux

func (self *Analyse) GetSeq(point int32, length int32, count int) (ret []Card) {
	for i := point - length + 1; i <= point; i++ {
		if len(self.Cards[i]) < count {
			return nil
		} else {
			ret = append(ret, self.Cards[i][:count]...)
		}
	}

	return
}

func (self *Analyse) GetJJ() (ret []Card) {
	if len(self.Cards[CardPoint_J1]) == 1 && len(self.Cards[CardPoint_J2]) == 1 {
		ret = append(ret, self.Cards[CardPoint_J1][0], self.Cards[CardPoint_J2][0])
	}

	return
}

func (self *Analyse) GetBombs() (ret [][]Card) {
	for _, v := range self.Cards {
		if len(v) == 4 {
			ret = append(ret, v)
		}
	}

	jj := self.GetJJ()
	if len(jj) > 0 {
		ret = append(ret, jj)
	}

	return
}
