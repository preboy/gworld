package chapter

import (
	"time"

	_ "core/log"
	"public/ec"
	"public/protocol/msg"
	"server/app"
	"server/battle"
	"server/config"
	"server/constant"
	"server/modules/drop"
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

	LootTs   int64    // 上次领取挂机奖励的时间
	BreakId  uint32   // 当前已完成的关卡ID
	Chapters []uint32 // 已领过奖励的章节ID
}

func (self *Chapter) Init(plr iPlayer) {
	self.plr = plr
}

func (self *Chapter) ToMsg() *msg.ChapterInfo {
	return &msg.ChapterInfo{
		LootTs:   self.LootTs,
		BreakId:  self.BreakId,
		Chapters: self.Chapters,
	}
}

// 拉取关卡信息
func (self *Chapter) ChapterInfo(req *msg.ChapterInfoRequest, res *msg.ChapterInfoResponse) {
	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 攻击关卡
func (self *Chapter) ChapterFighting(req *msg.ChapterFightingRequest, res *msg.ChapterFightingResponse) {
	conf := config.BreakConf.Query(self.BreakId + 1)
	if conf == nil {
		res.ErrorCode = ec.Conf_Invalid
		return
	}

	plrTrp := self.plr.CreateBattleTroop(req.Team)
	CreTrp := app.CreatureTeamToBattleTroop(conf.TeamId)

	if plrTrp == nil || CreTrp == nil {
		res.ErrorCode = ec.Failed
		return
	}

	b := battle.NewBattle(plrTrp, CreTrp)
	b.Calc()
	res.Win = b.GetResult()

	if res.Win {
		self.BreakId++

		proxy := app.NewItemProxy(constant.ItemLog_ChapterBreakPass)

		for _, v := range drop.Drop(self.plr, conf.LootId) {
			proxy.Add(v.Id, v.Cnt)
		}

		res.Rewards = proxy.Apply(self.plr).ToMsg()
	}

	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 领取章节奖励
func (self *Chapter) ChapterRewards(req *msg.ChapterRewardsRequest, res *msg.ChapterRewardsResponse) {
	conf := config.ChapterConf.Query(req.Id)
	if conf == nil {
		res.ErrorCode = ec.Failed // 未查到该章节
		return
	}

	if self.BreakId < conf.BreakEnd {
		res.ErrorCode = ec.Failed // 未完成该章节
		return
	}

	for _, v := range self.Chapters {
		if v == req.Id {
			res.ErrorCode = ec.Failed // 已领取过了
			return
		}
	}

	self.Chapters = append(self.Chapters, req.Id)

	proxy := app.NewItemProxy(constant.ItemLog_ChapterGot).SetArgs(req.Id)

	for _, v := range drop.Drop(self.plr, conf.DropId) {
		proxy.Add(v.Id, v.Cnt)
	}

	res.Rewards = proxy.Apply(self.plr).ToMsg()

	res.Info = self.ToMsg()
	res.ErrorCode = ec.OK
}

// 领取挂机奖励
func (self *Chapter) ChapterLoot(req *msg.ChapterLootRequest, res *msg.ChapterLootResponse) {
	now := time.Now().Unix()
	sec := now - self.LootTs
	if sec > 3600 { // 1个小时
		sec = 3600
	}

	if sec < 60 {
		res.ErrorCode = ec.Failed // 至少一分钟
		return
	}

	conf := config.BreakConf.Query(self.BreakId)
	if conf == nil {
		res.ErrorCode = ec.Conf_Invalid
		return
	}

	self.LootTs = now

	proxy := app.NewItemProxy(constant.ItemLog_ChapterLoot)

	// 此处关于性能问题，以后再优化
	for i := 0; i < int(sec)/60; i++ {
		for _, v := range drop.Drop(self.plr, conf.LootId) {
			proxy.Add(v.Id, v.Cnt)
		}
	}

	res.Rewards = proxy.Apply(self.plr).ToMsg()

	// 检测条件是否足够
	res.ErrorCode = ec.OK
}
