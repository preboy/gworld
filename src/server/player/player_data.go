package player

type PlayerData struct {
	// 这里的数据就是要存入DB的数据
	Name string `bson:nName"`
	Acct string `bson:"acct"`
	Pid  uint64 `bson:"pid"`
}

func GetPlayerData(acct string) *PlayerData {
	return nil
}

func CreatePlayerData(acct string) *PlayerData {
	return nil
}
