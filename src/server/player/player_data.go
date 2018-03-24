package player

import (
	"core/db"
	"core/log"
	"core/utils"
	"gopkg.in/mgo.v2"
	"server/db_mgr"
	"server/game"
	"time"
)

type TItemTimed map[uint32]uint64 // "20180226" => cnt 表示2018-02-26之后过期

type PlayerData struct {
	// 这里的数据就是要存入DB的数据
	Pid      uint64 `bson:"pid"`
	Acct     string `bson:"acct"`
	Name     string `bson:"name"`
	ShowName string `bson:"show_name"`
	PlatName string `bson:"plat_name"`
	ServerID uint32 `bson:"server_id"`

	// remark:  map的键必要是字符串  加载之后，写入之前需要特别处理
	Heros_bson      map[string]*game.Hero `bson:"heros"`       // 英雄
	Items_bson      map[string]uint64     `bson:"items"`       // 道具
	ItemsTimed_bson map[string]TItemTimed `bson:"items_timed"` // 限时道具
	Heros           map[uint32]*game.Hero `bson:"-"`
	Items           map[uint32]uint64     `bson:"-"`
	ItemsTimed      map[uint32]TItemTimed `bson:"-"`
	Level           uint32                `bson:"level"`       // 等级
	VipLevel        uint32                `bson:"vip_level"`   // VIP等级
	Last_update     int64                 `bson:"last_update"` // 最后一次处理数据的时间
	Male            bool                  `bson:"male"`        // 性别(默认:女)
	LoginTimes      uint32                `bson:"login_times"` // 登录次数
}

func (self *Player) GetData() *PlayerData {
	return self.data
}

func (self *PlayerData) SetName(name string) {
	self.ShowName = name
	self.Name = utils.U32toa(self.ServerID) + "." + name
}

func (self *Player) Save() {
	self.on_before_save()
	err := db_mgr.GetDB().UpsertByCond(
		db_mgr.Table_name_players,
		db.Condition{
			"acct": self.data.Acct,
		},
		self.data,
	)
	if err != nil {
		log.Error("Player.Save: Faild")
	}
}

func (self *Player) on_after_load() {
	self.last_update = self.data.Last_update

	data := self.GetData()

	data.Heros = make(map[uint32]*game.Hero)
	data.Items = make(map[uint32]uint64)
	data.ItemsTimed = make(map[uint32]TItemTimed)

	for k, v := range data.Heros_bson {
		key := utils.Atou32(k)
		data.Heros[key] = v
	}
	for k, v := range data.Items_bson {
		key := utils.Atou32(k)
		data.Items[key] = v
	}
	for k, v := range data.ItemsTimed_bson {
		key := utils.Atou32(k)
		data.ItemsTimed[key] = v
	}
}

func (self *Player) on_before_save() {
	self.data.Last_update = self.last_update

	data := self.GetData()

	data.Heros_bson = make(map[string]*game.Hero)
	data.Items_bson = make(map[string]uint64)
	data.ItemsTimed_bson = make(map[string]TItemTimed)

	for k, v := range data.Heros {
		key := utils.U32toa(k)
		data.Heros_bson[key] = v
	}
	for k, v := range data.Items {
		key := utils.U32toa(k)
		data.Items_bson[key] = v
	}
	for k, v := range data.ItemsTimed {
		key := utils.U32toa(k)
		data.ItemsTimed_bson[key] = v
	}
}

// ------------------ global ------------------

func LoadData() {
	// 加载DB中所有的玩家
	arr := []*PlayerData{}
	err := db_mgr.GetDB().GetAllObjects(
		db_mgr.Table_name_players,
		&arr,
	)
	if err != nil && err != mgo.ErrNotFound {
		log.Error("load all PlayerData err :", err)
		return
	}

	for _, data := range arr {
		plr := NewPlayer()
		plr.AssociateData(data)
	}

	log.Info("[%d] player loaded !", len(arr))
}

func SaveData() {
	// 所有玩家存盘
	// TODO
}

func CreatePlayerData(acct string) *PlayerData {
	pid := game.GeneralPlayerID()
	nam := game.GeneralPlayerName(pid)

	data := &PlayerData{
		Acct:        acct,
		Pid:         pid,
		ServerID:    game.GetServerConfig().Server_id,
		Last_update: time.Now().Unix() * 1000,
	}
	data.SetName(nam)

	return data
}
