package player

import (
	"time"

	"core/db"
	"core/log"
	// "core/utils"
	"game/app"
	"game/dbmgr"
	"game/modules/achv"
	"game/modules/chapter"
	"game/modules/quest"
)

type PlayerData struct {
	owner *Player

	// 这里的数据就是要存入DB的数据
	Pid   string `bson:"pid"`
	Acct  string `bson:"acct"`
	Name  string `bson:"name"`
	Plat  string `bson:"plat"`
	SvrId string `bson:"svrid"`

	// remark:  map的键必要是字符串  加载之后，写入之前需要特别处理
	Heros    map[uint32]*app.Hero `bson:"heros"`
	Items    map[uint32]uint64    `bson:"items"`
	Exp      uint64               `bson:"exp"`   // 经验
	Level    uint32               `bson:"lv"`    // 等级
	VipLevel uint32               `bson:"viplv"` // VIP等级
	Male     bool                 `bson:"male"`  // 性别(默认:女)

	CreateTs   time.Time `bson:"create_ts"`   // 创建角色时间
	LoginTs    time.Time `bson:"login_ts"`    // 最近登录时间
	LogoutTs   time.Time `bson:"logout_ts"`   // 最近下线时间
	LoginTimes uint32    `bson:"login_times"` // 总登录次数

	// modules data
	Growth  *achv.Growth     `bson:"growth"`
	Achv    *achv.Achv       `bson:"achv"`
	Quest   *quest.Quest     `bson:"quest"`
	Chapter *chapter.Chapter `bson:"chapter"`
}

// ============================================================================

func (self *PlayerData) Init(plr *Player) {
	self.owner = plr

	if self.Growth == nil {
		self.Growth = achv.NewGrowth()
	}
	self.Growth.Init(plr)

	if self.Achv == nil {
		self.Achv = achv.NewAchv()
	}
	self.Achv.Init(plr)

	if self.Quest == nil {
		self.Quest = quest.NewQuest()
	}
	self.Quest.Init(plr)

	if self.Chapter == nil {
		self.Chapter = chapter.NewChapter()
	}

	self.Chapter.Init(plr)
}

func (self *Player) Save() {
	err := dbmgr.GetDB().UpsertByCond(
		dbmgr.Table_name_player,
		db.Condition{
			"acct": self.data.Acct,
		},
		self.data,
	)
	if err != nil {
		log.Error("Player.Save: Faild")
	}
}

func (self *Player) AsyncSave() {
}

func (self *Player) on_after_load() {
	// data := self.GetData()

	// data.Heros = make(map[uint32]*app.Hero)
	// data.Items = make(map[uint32]uint64)

	// for k, v := range data.Heros_bson {
	// 	key := utils.Atou32(k)
	// 	data.Heros[key] = v
	// }
	// for k, v := range data.Items_bson {
	// 	key := utils.Atou32(k)
	// 	data.Items[key] = v
	// }
}

// ============================================================================
// player methond

func (self *Player) GetData() *PlayerData {
	return self.data
}

// ============================================================================
// exporter

func CreatePlayerData(acct string) *PlayerData {
	now := time.Now()
	pid, name := app.GeneralPlayerID()

	data := &PlayerData{
		Pid:      pid,
		Acct:     acct,
		Name:     name,
		Level:    1,
		SvrId:    app.GetGameId(),
		CreateTs: now,
	}

	// async save

	return data
}

func LoadPlayerDataFromDB(pid string) *PlayerData {
	// todo
	return nil
}

// ============================================================================
// data member export

func (self *Player) Getchapter() *chapter.Chapter {
	return self.data.Chapter
}
