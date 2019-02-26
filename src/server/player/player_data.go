package player

import (
	"time"

	"core/db"
	"core/log"
	"core/utils"
	"server/app"
	"server/db_mgr"
	"server/modules/achv"
	"server/modules/quest"
)

type TItemTimed map[uint32]uint64 // "20180226" => cnt 表示2018-02-26之后过期

type PlayerData struct {
	owner *Player

	// 这里的数据就是要存入DB的数据
	Pid      string `bson:"pid"`
	Acct     string `bson:"acct"`
	Name     string `bson:"name"`
	ShowName string `bson:"show_name"`
	PlatName string `bson:"plat_name"`
	ServerID uint32 `bson:"server_id"`

	// remark:  map的键必要是字符串  加载之后，写入之前需要特别处理
	Heros_bson      map[string]*app.Hero  `bson:"heros"`       // 英雄
	Items_bson      map[string]uint64     `bson:"items"`       // 道具
	ItemsTimed_bson map[string]TItemTimed `bson:"items_timed"` // 限时道具
	Heros           map[uint32]*app.Hero  `bson:"-"`
	Items           map[uint32]uint64     `bson:"-"`
	ItemsTimed      map[uint32]TItemTimed `bson:"-"`
	Exp             uint64                `bson:"exp"`         // 经验
	Level           uint32                `bson:"level"`       // 等级
	VipLevel        uint32                `bson:"vip_level"`   // VIP等级
	LastUpdate      int64                 `bson:"last_update"` // 最后一次处理数据的时间
	Male            bool                  `bson:"male"`        // 性别(默认:女)

	CreateTs   time.Time `bson:"create_ts"`   // 创建角色时间
	LoginTs    time.Time `bson:"login_ts"`    // 最近登录时间
	OfflineTs  time.Time `bson:"offline_ts"`  // 最近下线时间
	LoginTimes uint32    `bson:"login_times"` // 总登录次数

	// modules data
	Growth *achv.Growth `bson:"growth"`
	Achv   *achv.Achv   `bson:"achv"`
	Quest  *quest.Quest `bson:"quest"`
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

}

func (self *PlayerData) SetName(name string) {
	self.ShowName = name
	self.Name = utils.U32toa(self.ServerID) + "." + name
}

func (self *Player) Save() {
	self.on_before_save()
	err := db_mgr.GetDB().UpsertByCond(
		db_mgr.Table_name_players,
		db.Condition{
			"acct": self.data.Acct,
		},
		self.data,
	)
	if err != nil {
		log.Error("Player.Save: Faild")
	}
}

func (self *Player) on_after_load() {
	self.last_update = self.data.LastUpdate

	data := self.GetData()

	data.Heros = make(map[uint32]*app.Hero)
	data.Items = make(map[uint32]uint64)
	data.ItemsTimed = make(map[uint32]TItemTimed)

	for k, v := range data.Heros_bson {
		key := utils.Atou32(k)
		data.Heros[key] = v
	}
	for k, v := range data.Items_bson {
		key := utils.Atou32(k)
		data.Items[key] = v
	}
	for k, v := range data.ItemsTimed_bson {
		key := utils.Atou32(k)
		data.ItemsTimed[key] = v
	}
}

// ============================================================================
// player methond

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *Player) on_before_save() {
	self.data.LastUpdate = self.last_update

	data := self.GetData()

	data.Heros_bson = make(map[string]*app.Hero)
	data.Items_bson = make(map[string]uint64)
	data.ItemsTimed_bson = make(map[string]TItemTimed)

	for k, v := range data.Heros {
		key := utils.U32toa(k)
		data.Heros_bson[key] = v
	}
	for k, v := range data.Items {
		key := utils.U32toa(k)
		data.Items_bson[key] = v
	}
	for k, v := range data.ItemsTimed {
		key := utils.U32toa(k)
		data.ItemsTimed_bson[key] = v
	}
}

// ============================================================================
// exporter

func CreatePlayerData(acct string) *PlayerData {
	pid := app.GeneralPlayerID()
	nam := app.GeneralPlayerName(pid)
	now := time.Now()

	data := &PlayerData{
		Acct:       acct,
		Pid:        pid,
		Level:      1,
		ServerID:   app.GetServerConfig().Server_id,
		LastUpdate: now.Unix() * 1000,
		CreateTs:   now,
		LoginTs:    now,
	}

	data.SetName(nam)

	return data
}
