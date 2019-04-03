package app

import (
	"core/db"
	"core/log"
	"game/db_mgr"

	"os"
	"time"
)

var _sd *ServerData

type ServerData struct {
	ServerOpenTime int64  `bson:"server_open_time"` // 新服开启时间
	ServerSaveTime int64  `bson:"server_save_time"` // 数据最近保存时间
	IdSeq          uint32 `bson:"id_seq"`           // 玩家ID序列索引
	NewFlag        bool   `bson:"-"`                // 新服标识
	ServerId       uint32 `bson:"server_id"`        // 创建角色的服务器ID
}

// 加载服务器全局数据
func LoadServerData() {
	if _sd == nil {
		var data ServerData
		err := db_mgr.GetDB().GetObjectByCond(
			db_mgr.Table_name_server,
			db.Condition{
				"server_id": _sc.Server_id,
			},
			&data,
		)

		// 新开服
		if db.IsNotFound(err) {
			data.ServerOpenTime = time.Now().Unix()
			data.IdSeq = 1
			data.NewFlag = true
			data.ServerId = _sc.Server_id
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

	db_mgr.GetDB().UpsertByCond(
		db_mgr.Table_name_server,
		db.Condition{
			"server_id": _sd.ServerId,
		},
		_sd,
	)
}

func GetServerData() *ServerData {
	return _sd
}
