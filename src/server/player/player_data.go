package player

import (
	"server/game"
)

type PlayerData struct {
	// 这里的数据就是要存入DB的数据
	Name string `bson:nName"`
	Acct string `bson:"acct"`
	Pid  uint64 `bson:"pid"`
}

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *Player) Save() {

}

// --------- global

func GetPlayerData(acct string) *PlayerData {
	return nil
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
