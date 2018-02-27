package player

import (
	"core/db"
	"core/log"
	"server/db_mgr"
	"server/game"
)

type TItemTimed map[uint32]uint64 // "20180226" => cnt 表示2018-02-26之后过期

type PlayerData struct {
	// 这里的数据就是要存入DB的数据
	Name string `bson:nName"`
	Acct string `bson:"acct"`
	Pid  uint64 `bson:"pid"`

	Heros       map[uint32]*game.Hero `bson:"heros"`       // 英雄
	Items       map[uint32]uint64     `bson:"items"`       // 道具
	ItemsTimed  map[uint32]TItemTimed `bson:"items_timed"` // 限时道具
	Level       uint32                `bson:"level"`       // 等级
	VipLevel    uint32                `bson:"vip_level"`   // VIP等级
	Last_update int64                 `bson:"last_update"` // 最后一次处理数据的时间
	Male        bool                  `bson:"male"`        // 性别(默认:女)
	LoginTimes  uint32                `bson:"login_times"` // 登录次数
}

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *Player) Save() {
	self.data.Last_update = self.last_update
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

// ------------------ global ------------------

func LoadPlayerData(acct string) *PlayerData {
	var data PlayerData
	err := db_mgr.GetDB().GetObjectByCond(
		db_mgr.Table_name_players,
		db.Condition{
			"acct": acct,
		},
		&data,
	)

	if err != nil {
		return nil
	}

	return &data
}

func CreatePlayerData(acct string) *PlayerData {

	pid := game.GeneralPlayerID()
	nam := game.GeneralPlayerName(pid)

	data := &PlayerData{
		Acct: acct,
		Pid:  pid,
		Name: nam,
	}

	return data
}
