// respect to Mr.Long

package db

import (
	"core/log"
	"core/utils"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// ============================================================================

type M bson.M
type Change mgo.Change
type Condition bson.M

type Database struct {
	host      string       // connection string
	use_clone bool         // use cloned session ?
	session   *mgo.Session // master session
}

// ============================================================================

func NewDatabase() *Database {
	return &Database{}
}

// ============================================================================

func (self *Database) s_get() *mgo.Session {
	s := self.session

	if self.use_clone {
		s = s.Clone()
	}

	return s
}

func (self *Database) s_free(s *mgo.Session) {
	if self.use_clone {
		s.Close()
	}
}

func (self *Database) Open(host string, use_clone bool) {

	// create master session
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	s.SetMode(mgo.Strong, true)

	// params
	self.host = host
	self.use_clone = use_clone
	self.session = s
}

func (self *Database) Close() {
	self.session.Close()
}

func (self *Database) HasDB() bool {
	session := self.s_get()
	defer self.s_free(session)

	arr, err := session.DatabaseNames()
	if err != nil {
		log.Error("<HasDB>:", err)
		return false
	}

	name := session.DB("").Name

	for _, v := range arr {
		if v == name {
			return true
		}
	}

	return false
}

func (self *Database) HasCollection(coll string) bool {
	session := self.s_get()
	defer self.s_free(session)

	arr, err := session.DB("").CollectionNames()
	if err != nil {
		log.Error("<HasCollection>:", err)
		return false
	}

	for _, v := range arr {
		if v == coll {
			return true
		}
	}

	return false
}

func (self *Database) HasIndex(coll string, name string) bool {
	session := self.s_get()
	defer self.s_free(session)

	arr, err := session.DB("").C(name).Indexes()
	if err != nil {
		return false
	}

	for _, v := range arr {
		if v.Name == name {
			return true
		}
	}

	return false
}

func (self *Database) CreateCappedCollection(coll string, size int) {

	if self.HasCollection(coll) {
		return
	}

	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").Run(bson.D{
		{"create", coll},
		{"capped", true},
		{"size", size},
	}, nil)

	if err != nil {
		log.Error("<CreateCappedCollection>:", err)
	}
}

func (self *Database) CreateTTLIndex(coll string, name string, key string, sec int) {

	if self.HasIndex(coll, name) {
		return
	}

	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).EnsureIndex(mgo.Index{
		Name:        name,
		Key:         []string{key},
		ExpireAfter: time.Duration(sec) * time.Second,
	})
	if err != nil {
		log.Error("<CreateTTLIndex>: field: %s.%s, error: %s", coll, key, err)
	}
}

func (self *Database) CreateIndex(coll string, name string, keys []string, unique bool) {

	if self.HasIndex(coll, name) {
		return
	}

	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).EnsureIndex(mgo.Index{
		Name:   name,
		Key:    keys,
		Unique: unique,
	})
	if err != nil {
		log.Error("<CreateIndex>: field: %s.%v, error: %s", coll, keys, err)
	}
}

func (self *Database) EnableSharding() {

	if self.HasDB() {
		return
	}

	session := self.s_get()
	defer self.s_free(session)

	err := session.Run(bson.D{
		{"enableSharding", session.DB("").Name},
	}, nil)

	if err != nil {
		log.Error("<EnableSharding>:", err)
	}
}

func (self *Database) ShardCollection(coll string) {

	if self.HasCollection(coll) {
		return
	}

	session := self.s_get()
	defer self.s_free(session)

	err := session.Run(bson.D{
		{"shardCollection", fmt.Sprintf("%s.%s", session.DB("").Name, coll)},
		{"key", bson.M{"_id": "hashed"}},
	}, nil)

	if err != nil {
		log.Error("<ShardCollection>:", err)
	}
}

// ============================================================================

func (self *Database) GetObject(coll string, id interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).FindId(id).One(obj)
	if is_critical(session, err) {
		log.Error("db.GetObject():", err, coll, id)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetObjectByCond(coll string, cond interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Find(cond).One(obj)
	if is_critical(session, err) {
		log.Error("db.GetObjectByCond(): %v, %v, %v", err, coll, cond)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetProjection(coll string, id interface{}, proj interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).FindId(id).Select(proj).One(obj)
	if is_critical(session, err) {
		log.Error("db.GetProjection():", err, coll, id, proj)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetProjectionByCond(coll string, cond interface{}, proj interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Find(cond).Select(proj).One(obj)
	if is_critical(session, err) {
		log.Error("db.GetProjectionByCond():", err, coll, cond, proj)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetAllObjects(coll string, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Find(nil).All(obj)
	if is_critical(session, err) {
		log.Error("db.GetAllObjects():", err, coll)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetAllObjectsByCond(coll string, cond interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Find(cond).All(obj)
	if is_critical(session, err) {
		log.Error("db.GetAllObjectsByCond():", err, coll, cond)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) GetAllProjectionsByCond(coll string, cond interface{}, proj interface{}, obj interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Find(cond).Select(proj).All(obj)
	if is_critical(session, err) {
		log.Error("db.GetAllProjectionsByCond():", err, coll, cond, proj)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) Insert(coll string, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Insert(doc)
	if is_critical(session, err) {
		log.Error("db.Insert():", err, coll, doc)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) Remove(coll string, id interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).RemoveId(id)
	if is_critical(session, err) {
		log.Error("db.Remove():", err, coll, id)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) RemoveByCond(coll string, cond interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Remove(cond)
	if is_critical(session, err) {
		log.Error("db.RemoveByCond():", err, coll, cond)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) RemoveAll(coll string, cond interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	_, err := session.DB("").C(coll).RemoveAll(cond)
	if is_critical(session, err) {
		log.Error("db.RemoveAll():", err, coll, cond)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) Update(coll string, id interface{}, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).UpdateId(id, doc)
	if is_critical(session, err) {
		log.Error("db.Update():", err, coll, id)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) UpdateByCond(coll string, cond interface{}, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	err := session.DB("").C(coll).Update(cond, doc)
	if is_critical(session, err) {
		log.Error("db.UpdateByCond():", err, coll, cond, doc)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) UpdateAll(coll string, cond interface{}, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	_, err := session.DB("").C(coll).UpdateAll(cond, doc)
	if is_critical(session, err) {
		log.Error("db.UpdateAll():", err, coll, cond, doc)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) Upsert(coll string, id interface{}, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	_, err := session.DB("").C(coll).UpsertId(id, doc)
	if is_critical(session, err) {
		log.Error("db.Upsert():", err, coll, id)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) UpsertByCond(coll string, cond interface{}, doc interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	_, err := session.DB("").C(coll).Upsert(cond, doc)
	if is_critical(session, err) {
		log.Error("db.UpsertByCond():", err, coll, cond)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) FindAndModify(coll string, cond interface{}, chg Change, proj interface{}, r interface{}) error {
	session := self.s_get()
	defer self.s_free(session)

	_, err := session.DB("").C(coll).Find(cond).Select(proj).Apply(mgo.Change(chg), r)
	if is_critical(session, err) {
		log.Error("db.FindAndModify():", err, coll, cond, proj)
		log.Error(utils.Callstack())
	}
	return err
}

func (self *Database) Execute(f func(*mgo.Session)) {
	session := self.s_get()
	defer self.s_free(session)

	f(session)
}

// ============================================================================

func IsNotFound(err error) bool {
	return err == mgo.ErrNotFound
}

func IsDup(err error) bool {
	return mgo.IsDup(err)
}

// ============================================================================

func is_critical(session *mgo.Session, err error) bool {
	b := (err != nil && err != mgo.ErrNotFound && !mgo.IsDup(err))
	if b {
		session.Refresh()
	}
	return b
}
