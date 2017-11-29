package db

var _db *Database

func init() {
	_db = NewDatabase()
}

func GetDB() *Database {
	return _db
}

func Open(url string) bool {
	return _db.Open(url)
}

func Close() {
	_db.Close()
}
