package config

import (
	"core/log"
)

// ============================================================================

type Quest struct {
	Id      uint32       `json:"id"`
	Conds   []*QuestCond `json:"conds"`
	Type    uint32       `json:"type"`
	Head    string       `json:"head"`
	Desc    string       `json:"desc"`
	Tasks   []*QuestTask `json:"tasks"`
	Rewards []*ItemConf  `json:"rewards"`
}

type QuestTable struct {
	items map[uint32]*Quest
}

// ============================================================================

type QuestCond struct {
	Type uint32  `json:"type"`
	Args []int32 `json:"args"`
}

type QuestTask struct {
	Type uint32  `json:"type"`
	Text string  `json:"text"`
	Args []int32 `json:"args"`
}

// ============================================================================

var (
	QuestConf = &QuestTable{}
)

// ============================================================================

func (self *QuestTable) Load() bool {
	file := "Quest.json"
	var arr []*Quest

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Quest)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *QuestTable) Query(id uint32) *Quest {
	return self.items[id]
}

func (self *QuestTable) Items() map[uint32]*Quest {
	return self.items
}
