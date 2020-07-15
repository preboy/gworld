package player

import (
	"game/modules/achv"
	"game/modules/quest"
	"game/modules/task"
)

func (self *Player) GetId() string {
	return self.data.Pid
}

func (self *Player) GetLv() uint32 {
	return self.data.Lv
}

func (self *Player) GetName() string {
	return self.data.Name
}

func (self *Player) GetGrowth() *achv.Growth {
	return self.data.Growth
}

func (self *Player) GetAchv() *achv.Achv {
	return self.data.Achv
}

func (self *Player) GetQuest() *quest.Quest {
	return self.data.Quest
}

func (self *Player) GetTask() *task.Task {
	return self.data.Task
}
