package config

import (
	"core/log"
)

// ============================================================================

type Quest struct {
	Id       uint32       `json:"id"`
	Type     uint32       `json:"type"`
	Conds    []*QuestCond `json:"conds"`
	Title    string       `json:"title"`
	Desc     string       `json:"desc"`
	Commit   string       `json:"commit"`
	AcceptId uint32       `json:"accept_id"`
	CommitId uint32       `json:"commit_id"`
	Tasks    []*QuestTask `json:"tasks"`
	Rewards  []*ItemConf  `json:"rewards"`

	TaskMap map[uint32]*QuestTask
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
	Id         uint32          `json:"id"`
	Type       uint32          `json:"type"`
	Tip        string          `json:"tip"`
	Praise     string          `json:"praise"`
	NextId     uint32          `json:"next_id"`
	Save       bool            `json:"save"`
	TaskTalk   []*TalkOption   `json:"task_talk"`
	TaskKill   []*TaskTypeKill `json:"task_kill"`
	TaskGather []*ItemConf     `json:"task_gather"`
}

// ------------------------------------

// Npc谈话项
type TalkOption struct {
	Text   string `json:"text"`
	NextId uint32 `json:"next_id"`
}

type TaskTypeKill struct {
	Mid int32 `json:"mid"` // 要击杀的怪物ID(场景中的怪物)
	Cnt int32 `json:"cnt"` // 要击杀的次数
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
		v.TaskMap = make(map[uint32]*QuestTask)
		for _, t := range v.Tasks {
			v.TaskMap[t.Id] = t
		}

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
