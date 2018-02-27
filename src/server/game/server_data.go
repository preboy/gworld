package game

import (
	"core/db"
	"server/db_mgr"
)

import (
	"time"
)

var _sd *ServerData

type ServerData struct {
	Server_open_time int64  `bson:"server_open_time"` // 开服时间
	IdSeq            uint32 `bson:"id_seq"`           // 玩家ID序列索引
	NewFlag          bool   `bson:"-"`                // 新服标识
}

// 加载服务器全局数据
func LoadServerData() {
	if _sd == nil {
		var data ServerData
		err := db_mgr.GetDB().GetObjectByCond(
			db_mgr.Table_name_server,
			&data,
			db.Condition{
				"server_id": _sc.Server_id,
			},
		)

		// 新开服
		if err == nil {
			data.Server_open_time = time.Now().Unix()
			data.IdSeq = 1
			data.NewFlag = true
		}
		_sd = &data
	}
}

func SaveServerData() {

}

func GetServerData() *ServerData {
	return _sd
}
