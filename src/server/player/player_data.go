package player

import (
	"core/db"
	"core/log"
	"server/game"
)

type PlayerData struct {
	// 这里的数据就是要存入DB的数据
	Name string `bson:nName"`
	Acct string `bson:"acct"`
	Pid  uint64 `bson:"pid"`

	// 英雄
	Heros map[uint32]*game.Hero `bson:"heros"`
	// 背包
	Bag *PlayerBag `bson:"bag"`
}

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *Player) Save() {
	err := db.GetDB().Insert("player_data", self.data)
	if err != nil {
		log.Error("Faild to save")
	}
}

// ------------------ global ------------------

func GetPlayerData(acct string) *PlayerData {
	var data PlayerData
	err := db.GetDB().GetObjectByCond(
		"player_data",
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
