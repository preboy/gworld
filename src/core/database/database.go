package database

import (
	"gopkg.in/mgo.v2"
)

type Database struct {
	s *mgo.Session
}

func (self *Database) Open(url string) {
	session, err := mgo.Dial(url)
	if err == nil {
		self.s = session
	}
}

func (self *Database) Close() {
	if self.s != nil {
		self.s.Close()
		self.s = nil
	}
}

func (self *Database) Insert(col string, data interface{}) {

}

func (self *Database) Save(col string, data interface{}) {

}
