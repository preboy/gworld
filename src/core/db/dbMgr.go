package db

var _db *Database

func init() {
	_db = NewDatabase()
}

func GetDB() *Database {
	return _db
}

func Open(addr string) {
	_db.Open(addr, true)
}

func Close() {
	_db.Close()
}
