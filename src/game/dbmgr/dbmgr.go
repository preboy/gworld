package dbmgr

import (
	"core/db"
)

var (
	_db_game *db.Database
	_db_stat *db.Database
)

const (
	Table_name_server   = "server" // 服务器全局数据
	Table_name_player   = "player"
	Table_name_activity = "activity"
)

func init() {
	_db_game = db.NewDatabase()
	_db_stat = db.NewDatabase()
}

func GetDB() *db.Database {
	return _db_game
}

func GetStat() *db.Database {
	return _db_stat
}

func Open(game, stat string) {
	_db_game.Open(game, true)
	_db_stat.Open(stat, true)
}

func Close() {
	_db_game.Close()
	_db_stat.Close()
}
