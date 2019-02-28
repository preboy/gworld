package chapter

import (
	_ "core/log"
	"public/ec"
	"public/protocol/msg"
	"server/app"
	_ "server/config"
)

// ============================================================================
// regular

type iPlayer interface {
	app.IPlayer

	//
}

// ============================================================================

func NewChapter() *Chapter {
	return &Chapter{}
}

// ============================================================================

type Chapter struct {
	plr iPlayer

	LootTs   int64   // 上次领取挂机奖励的时间
	BreakId  uint32  // 当前已完成的关卡ID
	Chapters []int32 // 已领过奖励的章节ID
}

func (self *Chapter) Init(plr iPlayer) {
	self.plr = plr
}

func (self *Chapter) ToMsg() *msg.ChapterInfo {
	return &msg.ChapterInfo{}
}

// 拉取关卡信息
func (self *Chapter) ChapterInfo(req *msg.ChapterInfoRequest, res *msg.ChapterInfoResponse) {
	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 领取章节奖励
func (self *Chapter) ChapterFighting(req *msg.ChapterFightingRequest, res *msg.ChapterFightingResponse) {
	// TODO

	// 是否完成
	// 是否已领取

	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 攻击关卡
func (self *Chapter) ChapterRewards(req *msg.ChapterRewardsRequest, res *msg.ChapterRewardsResponse) {
	// TODO

	// 检测条件是否足够
	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 领取挂机奖励
func (self *Chapter) ChapterLoot(req *msg.ChapterLootRequest, res *msg.ChapterLootResponse) {
	// TODO

	// 检测条件是否足够
	res.ErrorCode = ec.OK
}
