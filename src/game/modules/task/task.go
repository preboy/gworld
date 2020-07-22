package task

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"core/event"
	"core/log"
	"game/app"
	"game/config"
	"game/constant"
)

// ============================================================================
// interface

type IPlayer interface {
	app.IPlayer

	GetTask() *Task
}

type ITask interface {
	NewTaskData() interface{}
}

// ============================================================================
// register

var (
	_tasks = make(map[uint32]ITask, 128)
)

func RegTask(id uint32, t ITask) {
	if _tasks[id] != nil {
		log.Fatal("REPEATED task id: %d", id)
	}

	_tasks[id] = t
}

// ============================================================================
// TaskItem

type TaskItem struct {
	Id      uint32
	StartTs uint32      // 开始时间
	Data    interface{} // 活动数据
	Over    bool        // 是否完成
	Draw    bool        // 是否领奖
}

func (self *TaskItem) GetId() uint32 {
	return self.Id
}

func (self *TaskItem) GetStartTs() uint32 {
	return self.StartTs
}

func (self *TaskItem) IsOver() bool {
	return self.Over
}

func (self *TaskItem) SetOver() {
	self.Over = true
}

// ============================================================================
// task_map_t

type task_map_t map[uint32]*TaskItem

// ============================================================================
// marshalling

func (self task_map_t) GetBSON() (interface{}, error) {
	var arr []*TaskItem

	for _, v := range self {
		arr = append(arr, v)
	}

	return arr, nil
}

func (self *task_map_t) SetBSON(raw bson.Raw) error {
	var arr []*TaskItem

	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(task_map_t)
	for _, v := range arr {
		(*self)[v.Id] = v
	}

	return nil
}

// ============================================================================
// Task

type Task struct {
	plr IPlayer

	Tasks task_map_t
}

func NewTask() *Task {
	return &Task{}
}

func (self *Task) Init(plr IPlayer) {
	self.plr = plr

	self.Tasks = make(task_map_t)
}

func (self *Task) Add(id uint32) bool {
	// exist
	if self.Tasks[id] != nil {
		return false
	}

	if _tasks[id] == nil {
		return false
	}

	task := &TaskItem{
		Id:      id,
		Data:    _tasks[id].NewTaskData(),
		StartTs: uint32(time.Now().Unix()),
	}

	self.Tasks[id] = task

	return true
}

func (self *Task) Del(id uint32) {
	delete(self.Tasks, id)
}

func (self *Task) Commit(id uint32) {
	task := self.Tasks[id]
	if task == nil {
		return
	}

	if !task.Over {
		return
	}

	if task.Draw {
		return
	}

	conf := config.TaskConf.Query(id)
	if conf == nil {
		return
	}

	{
		proxy := app.NewItemProxy(constant.ItemLog_TaskDraw)

		for _, v := range conf.Rewards {
			proxy.Add(v.Id, v.Cnt)
		}

		proxy.Apply(self.plr)
		task.Draw = true
	}

	self.Del(id)
}

func (self *Task) GetData(id uint32) interface{} {
	if task, ok := self.Tasks[id]; ok {
		return task.Data
	}

	return nil
}

func (self *Task) GetTaskItem(id uint32) *TaskItem {
	if task, ok := self.Tasks[id]; ok {
		return task
	}

	return nil
}

// ============================================================================
// init

func init() {
	// check impl
	event.On(constant.EVT_SYS_ConfigLoaded, func(evt uint32, args ...interface{}) {
		launch := args[0].(bool)
		if !launch {
			return
		}

		for _, conf := range config.TaskConf.Items() {
			if _tasks[conf.Id] == nil {
				log.Warning("NOT IMPL task: id = %v", conf.Id)
			}
		}
	})
}
