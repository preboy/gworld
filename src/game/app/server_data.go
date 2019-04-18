package app

import (
	"core/db"
	"core/log"
	"game/dbmgr"
	"os"
	"time"
)

var _sd *ServerData

type ServerData struct {
	ServerOpenTime int64    `bson:"open_ts"` // 新服开启时间
	ServerSaveTime int64    `bson:"save_ts"` // 数据最近保存时间
	IdSeq          uint32   `bson:"id_seq"`  // 玩家ID序列索引
	Svr            string   `bson:"svr"`     // 当前服ID
	SvrSet         []string `bson:"svr_set"` // 被合入的服ID
}

// 加载服务器全局数据
func LoadServerData() {
	if _sd == nil {
		var data ServerData
		err := dbmgr.GetDB().GetObject(
			dbmgr.Table_name_server,
			1,
			&data,
		)

		gameid := GetGameId()

		// 新开服
		if db.IsNotFound(err) {
			data.ServerOpenTime = time.Now().Unix()
			data.IdSeq = 1
			data.Svr = gameid
		} else if err != nil {
			log.Println("LoadServerData: Error")
			os.Exit(-1)
		}

		if data.Svr != gameid {
			log.Println("gameid MISMATCHING", gameid, data.Svr)
			os.Exit(-2)
		}

		if time.Now().Unix() < data.ServerSaveTime {
			log.Println("DATE ROLL BACK, It's fatal")
			os.Exit(-3)
		}

		_sd = &data
	}
}

func SaveServerData() {
	if _sd == nil {
		return
	}

	_sd.ServerSaveTime = time.Now().Unix()

	dbmgr.GetDB().Upsert(
		dbmgr.Table_name_server,
		1,
		_sd,
	)
}

func GetServerData() *ServerData {
	return _sd
}
