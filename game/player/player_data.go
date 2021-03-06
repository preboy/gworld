package player

import (
	"time"

	"gworld/core/db"
	"gworld/core/log"
	"gworld/core/utils"

	"gworld/game/app"
	"gworld/game/dbmgr"
	"gworld/game/modules/achv"
	"gworld/game/modules/chapter"
	"gworld/game/modules/quest"
	"gworld/game/modules/task"

	"gworld/public/protocol/msg"
)

type PlayerInfo struct {
	Acct string
	Name string
	Key  string
	Svr  string
	SDK  string
	Pid  string
	Lv   uint32
}

type PlayerData struct {
	owner *Player

	// 这里的数据就是要存入DB的数据
	Pid  string `bson:"pid"`
	Key  string `bson:"key"`
	Name string `bson:"name"`
	Acct string `bson:"acct"`
	Plat string `bson:"plat"`
	Svr  string `bson:"svr"`
	SDK  string `bson:"sdk"`

	Exp  uint64 `bson:"exp"`  // 经验
	Lv   uint32 `bson:"lv"`   // 等级
	Vip  uint32 `bson:"vip"`  // VIP等级
	Male bool   `bson:"male"` // 性别(默认:女)

	Heros     hero_map_t     `bson:"heros"`
	Items     item_map_t     `bson:"items"`
	Instances instance_map_t `bson:"instances"`

	CreateTs   time.Time `bson:"create_ts"`   // 创建角色时间
	LoginTs    time.Time `bson:"login_ts"`    // 最近登录时间
	LogoutTs   time.Time `bson:"logout_ts"`   // 最近下线时间
	LoginTimes uint32    `bson:"login_times"` // 总登录次数

	// modules data
	Growth  *achv.Growth     `bson:"growth"`
	Achv    *achv.Achv       `bson:"achv"`
	Quest   *quest.Quest     `bson:"quest"`
	Chapter *chapter.Chapter `bson:"chapter"`
	Task    *task.Task       `bson:"task"`
}

// ============================================================================

func (self *PlayerData) Init(plr *Player) {
	self.owner = plr

	if self.Heros == nil {
		self.Heros = make(hero_map_t)
	}
	if self.Items == nil {
		self.Items = make(item_map_t)
	}
	if self.Instances == nil {
		self.Instances = make(instance_map_t)
	}

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

	if self.Task == nil {
		self.Task = task.NewTask()
	}
	self.Task.Init(plr)
}

func (self *PlayerData) to_player_info() *PlayerInfo {
	return &PlayerInfo{
		Acct: self.Acct,
		Name: self.Name,
		Key:  self.Key,
		Svr:  self.Svr,
		SDK:  self.SDK,
		Pid:  self.Pid,
		Lv:   self.Lv,
	}
}

func (self *PlayerData) ToMsg() *msg.PlayerDataResponse {
	res := &msg.PlayerDataResponse{}

	res.Acct = self.Acct
	res.Name = self.Name
	res.Pid = self.Pid
	res.Sid = self.owner.sid
	res.Exp = self.Exp
	res.Lv = self.Lv
	res.Vip = self.Vip
	res.Male = self.Male

	for id, cnt := range self.Items {
		res.Items = append(res.Items, &msg.Item{
			Flag: 0,
			Id:   id,
			Cnt:  int64(cnt),
		})
	}

	for _, hero := range self.Heros {
		res.Heros = append(res.Heros, hero.ToMsg())
	}

	return res
}

// ============================================================================

func (self *Player) Save() {
	// player data
	{
		err := dbmgr.GetDB().UpsertByCond(
			dbmgr.Table_name_player,
			db.M{
				"_id": self.data.Pid,
				"key": self.data.Key,
			},
			self.data,
		)

		if err != nil {
			log.Error("player data save FAILED")
		}
	}

	// player info
	{
		info := self.data.to_player_info()

		err := dbmgr.GetCenter().UpsertByCond(
			dbmgr.Table_name_player_info,
			db.M{
				"_id": info.Pid,
				"key": info.Key,
			},
			info,
		)

		if err != nil {
			log.Error("player info save FAILED")
		}
	}
}

func (self *Player) AsyncSave() {
	data := utils.CloneBsonObject(self.data)
	info := self.data.to_player_info()

	go func() {
		err := dbmgr.GetDB().UpsertByCond(
			dbmgr.Table_name_player,
			db.M{
				"_id": self.data.Pid,
				"key": self.data.Key,
			},
			data,
		)

		if err != nil {
			log.Error("player data async save FAILED")
		}
	}()

	go func() {
		err := dbmgr.GetCenter().UpsertByCond(
			dbmgr.Table_name_player_info,
			db.M{
				"_id": info.Pid,
				"key": info.Key,
			},
			info,
		)

		if err != nil {
			log.Error("player info async save FAILED")
		}
	}()
}

// ============================================================================
// player methond

func (self *Player) GetData() *PlayerData {
	return self.data
}

// ============================================================================
// exporter

func GetPlayerData(key, acct, svr, sdk string) *PlayerData {
	data := load_player_data(key)

	if data == nil {
		data = create_player_data(key, acct, svr, sdk)
	}

	return data
}

func load_player_data(key string) *PlayerData {
	var data PlayerData
	err := dbmgr.GetDB().GetObjectByCond(
		dbmgr.Table_name_player,
		db.M{
			"key": key,
		},
		&data,
	)

	if err != nil {
		log.Error("Loading PlayerData err: %v", err)
		return nil
	}

	return &data
}

func create_player_data(key, acct, svr, sdk string) *PlayerData {
	now := time.Now()

	pid, name := app.GeneralPlayerID()

	data := &PlayerData{
		Pid:      pid,
		Key:      key,
		Name:     name,
		Acct:     acct,
		Plat:     app.GetPlat(),
		Svr:      svr,
		SDK:      sdk,
		Lv:       1,
		CreateTs: now,
	}

	return data
}

// ============================================================================
// data member export

func (self *Player) GetChapter() *chapter.Chapter {
	return self.data.Chapter
}
