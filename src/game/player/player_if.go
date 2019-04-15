package player

import (
	"game/modules/quest"
)

func (self *Player) GetQuest() *quest.Quest {
	return self.data.Quest
}
