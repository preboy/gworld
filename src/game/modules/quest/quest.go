package quest

import (
	"core/event"
	"core/log"
	"game/app"
	"game/config"
	"game/constant"
	"public/ec"
	"public/protocol/msg"
)

// ============================================================================
// regular

type iPlayer interface {
	app.IPlayer

	//
}

// ============================================================================

const (
	QuestType_Main    uint32 = iota + 10 // 主线
	QuestType_Fortune                    // 奇遇
)

const (
	CondiType_Lv        uint32 = iota + 20 // 等级
	CondiType_Sex                          // 性别
	CondiType_DailyDur                     // 每日时间段内
	CondiType_WeeklyDur                    // 每周时间段内
)

const (
	TaskType_Talk uint32 = iota + 30
	TaskType_Kill
	TaskType_Gather
)

const (
	CommandType_Talk uint32 = iota + 40
	CommandType_Kill
	CommandType_Gather
)

// ============================================================================

type quest_item_t struct {
	Id   uint32       // 任务ID
	Task uint32       // 当前的task项   0:表示已完成所有的task项
	Data app.KV_map_t // 任务项数据
}

func (self *quest_item_t) to_msg() *msg.QuestInfo {
	m := &msg.QuestInfo{
		Id:   self.Id,
		Task: self.Task,
	}

	for k, v := range self.Data {
		m.Data = append(m.Data, &msg.QuestData{
			Key: k,
			Val: v,
		})
	}

	return m
}

// ============================================================================

func NewQuest() *Quest {
	return &Quest{}
}

func init() {
	event.On(constant.Evt_Plr_KillMonster, func(evt uint32, args ...interface{}) {
		plr := args[0].(iPlayer)
		mid := args[1].(int32)

		plr.GetData().Quest.OnKill(mid)
	})
}

// ============================================================================

type Quest struct {
	plr iPlayer

	LastId    uint32        // 上一个完成的任务(主线)
	Main      *quest_item_t // 主线任务
	Forture   *quest_item_t // 奇遇任务
	Selection app.KV_map_t  // 任务值
}

func (self *Quest) Init(plr iPlayer) {
	self.plr = plr

	if self.Selection == nil {
		self.Selection = make(app.KV_map_t)
	}
}

func (self *Quest) ToMsg(id uint32) *msg.QuestInfo {
	if self.Main != nil && self.Main.Id == id {
		return self.Main.to_msg()
	}

	if self.Forture != nil && self.Forture.Id == id {
		return self.Forture.to_msg()
	}

	return nil
}

func (self *Quest) ToMsgs() (ret []*msg.QuestInfo) {
	if self.Main != nil {
		ret = append(ret, self.Main.to_msg())
	}

	if self.Forture != nil {
		ret = append(ret, self.Forture.to_msg())
	}

	return
}

// 接任务
func (self *Quest) Accept(id uint32) uint32 {
	conf := config.QuestConf.Query(id)
	if conf == nil {
		return ec.Conf_Invalid
	}

	// TODO 检测条件是否满足
	if false {
		return ec.QUEST_Cond_Dissatisfy
	}

	q := &quest_item_t{
		Id:   id,
		Task: 1,
		Data: make(app.KV_map_t),
	}

	switch conf.Type {
	case QuestType_Main:
		{
			if id <= self.LastId {
				return ec.QUEST_Pass_Over
			}
			if self.Main != nil {
				return ec.QUEST_Not_Finish
			}
			self.Main = q
		}
	case QuestType_Fortune:
		{
			if self.Forture != nil {
				return ec.QUEST_Not_Finish
			}
			self.Forture = q
		}
	default:
		{
			log.Error("Unknown Quest Type", id)
		}
	}

	return ec.OK
}

// 提交任务
func (self *Quest) Commit(id uint32, r uint32) uint32 {
	var q *quest_item_t

	if self.Main != nil && self.Main.Id == id {
		q = self.Main
	} else if self.Forture != nil && self.Forture.Id == id {
		q = self.Forture
	}

	if q == nil {
		return ec.Failed
	}

	if q.Task == 0 {
		return ec.QUEST_Tasks_Over
	}

	conf := config.QuestConf.Query(id)
	task := conf.TaskMap[q.Task]
	next := uint32(0)

	switch task.Type {
	case TaskType_Talk:
		{
			if r < 1 || int(r) > len(task.TaskTalk) {
				return ec.QUEST_Task_Invalid_r
			}
			if task.Save {
				self.Selection[int32(id)] = int32(r)
			}

			next = task.TaskTalk[r-1].NextId
		}
	case TaskType_Kill:
		{ // 检测是否杀够
			for _, v := range task.TaskKill {
				if q.Data[v.Mid] < v.Cnt {
					return ec.QUEST_Task_Invalid_Kill
				}
			}

			next = task.NextId
		}
	case TaskType_Gather:
		{ // 检测包裹是否收集完成
			proxy := app.NewItemProxy(constant.ItemLog_QuestCommit).SetArgs(id)
			for _, v := range task.TaskGather {
				proxy.Sub(v.Id, v.Cnt)
			}
			if !proxy.Enough(self.plr) {
				return ec.Item_Not_Enough
			}
			proxy.Apply(self.plr)

			next = task.NextId
		}
	}

	{
		if next == 0 {
			next = q.Task + 1
		}

		if conf.TaskMap[next] == nil {
			next = 0
		}

		q.Task = next

		if next != 0 {
			// 初始化
		}
	}

	return ec.OK
}

// 完成任务
func (self *Quest) Finish(id uint32) uint32 {
	var q *quest_item_t
	var qt uint32 = QuestType_Main

	if self.Main != nil && self.Main.Id == id {
		q = self.Main
		qt = QuestType_Main
	} else if self.Forture != nil && self.Forture.Id == id {
		q = self.Forture
		qt = QuestType_Fortune
	}

	if q == nil {
		return ec.Failed
	}

	if q.Task != 0 {
		return ec.QUEST_Task_Unfinished
	}

	conf := config.QuestConf.Query(id)

	// 发放奖励，设置标识
	proxy := app.NewItemProxy(constant.ItemLog_QuestFinish).SetArgs(id)

	for _, v := range conf.Rewards {
		proxy.Add(v.Id, v.Cnt)
	}

	proxy.Apply(self.plr)

	// 任务完成
	if qt == QuestType_Main {
		self.Main = nil
	} else if qt == QuestType_Fortune {
		self.Forture = nil
	}

	return ec.OK
}

// 放弃任务
func (self *Quest) Cancel(id uint32) uint32 {
	if self.Main != nil && self.Main.Id == id {
		self.Main = nil
	} else if self.Forture != nil && self.Forture.Id == id {
		self.Forture = nil
	}

	// ToMsg
	return ec.OK
}

func (self *Quest) OnKill(mid int32) {

	var check = func(q *quest_item_t) {
		if q == nil {
			return
		}

		if q.Task == 0 {
			return
		}

		conf := config.QuestConf.Query(q.Id)
		task := conf.TaskMap[q.Task]

		if task.Type != TaskType_Kill {
			return
		}

		for _, v := range task.TaskKill {
			if v.Mid == mid {
				if q.Data[mid] < v.Cnt {
					q.Data[mid]++
				}
			}
		}
	}

	check(self.Main)
	check(self.Forture)
}
