package task

import (
	"game/app"
	"game/config"

	"time"
)

// ============================================================================
// if

type iPlayer interface {
	app.IPlayer

	GetTask() *Task
}

// ============================================================================
// task item

type task_item_t struct {
	Id      uint32 // 原型ID
	StartTs uint32 // 开始时间
	Data    app.KV_map_t
	Finish  bool
}

func (self *task_item_t) GetVal(id int32) int32 {
	return self.Data[id]
}

func (self *task_item_t) SetVal(id int32, val int32) {
	self.Data[id] = val
}

// ============================================================================
// Task

type Task struct {
	plr iPlayer

	Tasks map[uint32]*task_item_t
}

func NewTask() *Task {
	return &Task{}
}

func (self *Task) Init(plr iPlayer) {
	self.plr = plr

	self.Tasks = make(map[uint32]*task_item_t)
}

func (self *Task) Add(id uint32) bool {

	conf := config.TaskConf.Query(id)
	if conf == nil {
		return false
	}

	if self.Tasks[id] != nil {
		return false
	}

	self.Tasks[id] = &task_item_t{
		Id:      id,
		StartTs: uint32(time.Now().Unix()),
		Data:    make(app.KV_map_t),
		Finish:  false,
	}

	return true
}

func (self *Task) Del(id uint32) {
	delete(self.Tasks, id)
}

func (self *Task) Get(id uint32) *task_item_t {
	return self.Tasks[id]
}

func (self *Task) Commit(id uint32) {
	//	todo 发奖励

	conf := config.TaskConf.Query(id)
	if conf == nil {
		return
	}

	if self.Tasks[id] != nil {
		return
	}

	// add .

	self.Del(id)
}
