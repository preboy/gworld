package db_mgr

import (
	"core/db"
)

var _db *db.Database

const (
	Table_name_server  = "server"  // 服务器全局数据
	Table_name_players = "players" // 玩家表
)

func init() {
	_db = db.NewDatabase()
}

func GetDB() *db.Database {
	return _db
}

func Open(addr string) {
	_db.Open(addr, true)
}

func Close() {
	_db.Close()
}
