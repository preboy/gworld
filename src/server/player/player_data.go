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

	Heros      map[uint32]*game.Hero `bson:"heros"`       // 英雄
	Items      map[uint32]uint64     `bson:"items"`       // 道具
	ItemsTimed map[uint32]TItemTimed `bson:"items_timed"` // 限时道具
	VipLevel   uint32                `bson:"vip_level"`   // VIP 等级
}

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *Player) Save() {
	err := db_mgr.GetDB().Insert(db_mgr.Table_name_players, self.data)
	if err != nil {
		log.Error("Faild to save")
	}
}

// ------------------ global ------------------

func GetPlayerData(acct string) *PlayerData {
	var data PlayerData
	err := db_mgr.GetDB().GetObjectByCond(
		db_mgr.Table_name_players,
		&data,
		db.Condition{
			"acct": acct,
		},
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
