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
	ServerOpenTime int64  `bson:"server_open_time"` // 新服开启时间
	ServerSaveTime int64  `bson:"server_save_time"` // 数据最近保存时间
	IdSeq          uint32 `bson:"id_seq"`           // 玩家ID序列索引
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

		// 新开服
		if db.IsNotFound(err) {
			data.ServerOpenTime = time.Now().Unix()
			data.IdSeq = 1
		} else if err != nil {
			log.Error("LoadServerData: Error")
		}

		if time.Now().Unix() < data.ServerSaveTime {
			log.Error("时间回退了吧，这样要不得的")
			os.Exit(-1)
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
