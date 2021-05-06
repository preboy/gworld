package lobby

type CardsType int

const (
	CardsTypeNIL              CardsType = iota + 0 // 非法
	CardsTypeA                                     // 单牌
	CardsTypeAA                                    // 一对
	CardsTypeJJ                                    // 王炸
	CardsTypeAAA                                   // 三不带
	CardsTypeAAAX                                  // 三带单
	CardsTypeAAAXX                                 // 三带对
	CardsTypeAAAA                                  // 炸弹
	CardsTypeAAAAXY                                // 四带二
	CardsTypeAAAAXXYY                              // 四带二对
	CardsTypeABCDE                                 // 顺子(大于5)
	CardsTypeAABBCC                                // 顺对
	CardsTypeAAABBB                                // 飞机
	CardsTypeAAABBBXY                              // 飞机带单
	CardsTypeAAABBBXXYY                            // 飞机带对
	CardsTypeAAAABBBB                              // 4顺
	CardsTypeAAAABBBBXYZM                          // 4顺带单
	CardsTypeAAAABBBBXXYYZZMM                      // 4顺带对
)

func cards_from_int32(cards []int32) (ret []Card) {
	for _, v := range cards {
		c := NewCard(v)
		ret = append(ret, c)

		if !c.Valid() {
			panic("INVALID card")
		}
	}

	return
}

func cards_to_int32(cards []Card) (ret []int32) {
	for _, c := range cards {
		ret = append(ret, int32(c))
	}

	return
}

func get_cards_type(arr []int32) CardsType {
	cards := cards_from_int32(arr)

	switch len(cards) {

	case 1: // CardsTypeA

	case 2: // CardsTypeAA CardsTypeJJ

	case 3: // CardsTypeAAA

	case 4: // CardsTypeAAAA CardsTypeAAAB

	case 5: // CardsTypeAAAXX CardsTypeABCDE

	case 6: // CardsTypeAABBCC CardsTypeAAABBB CardsTypeAAAAXY CardsTypeABCDE

	case 7: // CardsTypeABCDE

	case 8: // CardsTypeAABBCC CardsTypeAAABBBXY CardsTypeAAAABBBB  CardsTypeABCDE

	case 9: // CardsTypeAAABBB CardsTypeABCDE

	case 10: // CardsTypeAABBCC CardsTypeAAABBBXXYY CardsTypeABCDE

	case 11: // CardsTypeABCDE

	case 12: // CardsTypeAABBCC CardsTypeAAABBBXY CardsTypeAAAABBBBXYZM CardsTypeABCDE

	case 13: // 莫法
		break

	case 14: // CardsTypeAABBCC

	case 15: // CardsTypeAAABBB CardsTypeAAABBBXXYY

	default:
		break
	}

	return CardsTypeNIL
}
