package app

import (
	"core/db"
	"core/log"
	"game/dbmgr"

	"time"
)

var _sd *ServerData

type ServerData struct {
	ServerOpenTime time.Time `bson:"open_ts"`      // 新服开启时间
	ServerSaveTime time.Time `bson:"save_ts"`      // 数据最近保存时间
	SeqId          uint32    `bson:"seq_id"`       // 玩家ID序列索引
	Svr            string    `bson:"svr"`          // 当前服ID
	SvrSet         []string  `bson:"svr_set"`      // 被合入的服ID
	SeqInstance    uint32    `bson:"seq_instance"` // 除玩家外的其它ID生成器
}

// 加载服务器全局数据
func LoadServerData() bool {
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
			data.ServerOpenTime = time.Now()
			data.SeqId = 1
			data.Svr = gameid
			data.SeqInstance = 1
		} else if err != nil {
			log.Error("LoadServerData: Error")
			return false
		}

		if data.Svr != gameid {
			log.Error("boot gameid = %s, db gameid = %s", gameid, data.Svr)
			return false
		}

		if time.Now().Before(data.ServerSaveTime) {
			log.Error("DATE ROLL BACK, It's fatal")
			return false
		}

		_sd = &data
	}

	return true
}

func SaveServerData() {
	if _sd == nil {
		return
	}

	_sd.ServerSaveTime = time.Now()

	dbmgr.GetDB().Upsert(
		dbmgr.Table_name_server,
		1,
		_sd,
	)
}

func GetServerData() *ServerData {
	return _sd
}

func IsValidGameId(gameid string) bool {
	if gameid == _sd.Svr {
		return true
	}

	for _, v := range _sd.SvrSet {
		if gameid == v {
			return true
		}
	}

	return false
}

// ----------------------------------------------------------------------------

func (self *ServerData) GetSeqId() uint32 {
	self.SeqId++
	return self.SeqId
}

func (self *ServerData) GetSeqInstance() uint32 {
	self.SeqInstance++
	return self.SeqInstance
}
