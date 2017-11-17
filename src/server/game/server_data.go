package game

var _sd *ServerData

type ServerData struct {
	Server_open_time int64  `bson:"server_open_time"` // 开服时间
	IdSeq            uint32 `bson:"id_seq"`           // 玩家ID序列索引
}

// 加载服务器全局数据
func LoadServerData() {
	if _sd == nil {
		sd := ServerData{}

		// load from db

		_sd = &sd
	}
}

func SaveServerData() {
}

func GetServerData() *ServerData {
	return _sd
}
