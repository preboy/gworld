package quest

import (
	"core/log"
	"core/utils"
	"gopkg.in/mgo.v2/bson"
	"public/ec"
	"public/protocol"
	"server/app"
	"server/config"
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
	Id     uint32           // 任务ID
	Task   uint32           // 当前的task项   0:表示已完成所有的task项
	Finish bool             // true表示已领取奖励
	Data   map[string]int32 // 任务项数据
}

func (self *quest_item_t) init() {
}

// ============================================================================

type quest_value_t map[int32]int32

func (self quest_value_t) GetBSON() (interface{}, error) {
	type QV struct {
		Id, Val int32
	}

	var arr []*QV
	for k, v := range self {
		arr = append(arr, &QV{Id: k, Val: v})
	}

	return arr, nil
}

func (self *quest_value_t) SetBSON(raw bson.Raw) error {
	type QV struct {
		Id, Val int32
	}

	var arr []*QV
	err := raw.Unmarshal(&arr)
	if err != nil {
		return err
	}

	*self = make(quest_value_t)
	for _, v := range arr {
		(*self)[v.Id] = v.Val
	}

	return nil
}

// ============================================================================

func NewQuest() *Quest {
	return &Quest{}
}

// ============================================================================

type Quest struct {
	plr iPlayer

	LastId    uint32        // 上一个完成的任务(主线)
	Main      *quest_item_t // 主线任务
	Forture   *quest_item_t // 奇遇任务
	Selection quest_value_t // 任务值
}

func (self *Quest) Init(plr iPlayer) {
	self.plr = plr

	if self.Selection == nil {
		self.Selection = make(quest_value_t)
	}
}

// 接任务
func (self *Quest) Accept(id uint32) int {
	conf := config.QuestConf.Query(id)
	if conf == nil {
		return ec.Conf_Invalid
	}

	// 检测条件是否满足
	if false {
		return ec.QUEST_Cond_Dissatisfy
	}

	q := &quest_item_t{
		Id:   id,
		Task: 1,
		Data: make(map[string]int32),
	}
	q.init()

	switch conf.Type {
	case QuestType_Main:
		{
			if id <= self.LastId {
				return ec.QUEST_Pass_Over
			}
			if self.Main != nil && !self.Main.Finish {
				return ec.QUEST_Not_Finish
			}
			self.Main = q
		}
	case QuestType_Fortune:
		{
			if self.Forture != nil && !self.Forture.Finish {
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
func (self *Quest) Commit(id uint32, r int32) int {
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
				self.Selection[int32(id)] = r
			}

			next = task.TaskTalk[r-1].NextId
		}
	case TaskType_Kill:
		{ // 检测是否杀够
			for _, v := range task.TaskKill {
				if q.Data[utils.I32toa(v.Mid)] < v.Cnt {
					return ec.QUEST_Task_Invalid_Kill
				}
			}

			next = task.NextId
		}
	case TaskType_Gather:
		{ // 检测包裹是否收集完成
			proxy := app.NewItemProxy(1) // todo
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

	if next == 0 {
		next = q.Task + 1
	}

	// 新的task

	return ec.OK
}

// 完成任务
func (self *Quest) Finish(id uint32) int {
	var q *quest_item_t

	if self.Main != nil && self.Main.Id == id {
		q = self.Main
	} else if self.Forture != nil && self.Forture.Id == id {
		q = self.Forture
	}

	if q == nil {
		return ec.Failed
	}

	if q.Finish {
		return ec.QUEST_Finish_Yet
	}

	if q.Task != 0 {
		return ec.QUEST_Task_Unfinished
	}

	conf := config.QuestConf.Query(id)

	// 发放奖励，设置标识
	proxy := app.NewItemProxy(protocol.MSG_CS_MarketBuy)
	// TODO

	for _, v := range conf.Rewards {
		proxy.Add(v.Id, v.Cnt)
	}

	proxy.Apply(self.plr)
	q.Finish = true

	return ec.OK
}

// 放弃任务
func (self *Quest) Cancel(id uint32) {

}
