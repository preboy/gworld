package config

import (
	"core/log"
)

// ============================================================================

type Task struct {
	Id      uint32      `json:"id"`
	Dur     uint32      `json:"dur"`
	Type    uint32      `json:"type"`
	Params  []int32     `json:"params"`
	Rewards []*ItemConf `json:"rewards"`
	Title   string      `json:"title"`
	Desc    string      `json:"desc"`
}

type TaskTable struct {
	items map[uint32]*Task
}

// ============================================================================

type TaskCond struct {
	Type uint32  `json:"type"`
	Args []int32 `json:"args"`
}

// ============================================================================

var (
	TaskConf = &TaskTable{}
)

// ============================================================================

func (self *TaskTable) Load() bool {
	file := "Task.json"
	var arr []*Task

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Task)
	for _, v := range arr {
		self.items[v.Id] = v
	}

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *TaskTable) Query(id uint32) *Task {
	return self.items[id]
}

func (self *TaskTable) Items() map[uint32]*Task {
	return self.items
}
