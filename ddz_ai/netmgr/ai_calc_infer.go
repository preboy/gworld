package netmgr

// 计算玩家手中可能的牌
func (self *AILogic) infer() {
	if len(self.rounds) == 0 {
		return
	}

	// 根据场下剩余的牌、他已出的牌、对手出的牌
	// 只推算可能有的牌

	for _, r := range self.rounds {

		for _, h := range r.hands {

			// me
			if h.pos == self.pos {
				continue
			}

			// pass
			if h.cards == nil {
				continue
			}

			//

		}
	}
}
