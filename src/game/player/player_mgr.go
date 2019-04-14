package player

import (
	"fmt"

	"core/log"
	"game/app"
	"game/constant"
	"core/event"
	"game/db_mgr"
	"gopkg.in/mgo.v2"
)

const (
	MAX_PLAYER_COUNT = 0X4000
)

var (
	_index int = 1
)

/*
服务器启动时从DB中加载一定数量的玩家到内存，存入_plrs_pid、_plrs_name
新玩家立即存盘，并加载到内存
登录期间的玩家存入_plrs_online
*/
var (
	// 在内存中的玩家，包括主动上线、从DB中被拉上线的
	_plrs_sid    = [MAX_PLAYER_COUNT]*Player{}                // 内存中的玩家
	_plrs_pid    = make(map[string]*Player, MAX_PLAYER_COUNT) // pid(all players)
	_plrs_name   = make(map[string]*Player, MAX_PLAYER_COUNT) // name
	_plrs_online = make(map[string]*Player, MAX_PLAYER_COUNT) // pid(在线的玩家)
)

func (self *Player) SetName(name string) {
	// TODO: checking valid & repeat

	old_name := self.data.Name
	new_name := fmt.Sprintf("%s-%s", app.GetGameId(), name)

	self.data.Name = new_name

	_plrs_name[old_name] = nil
	_plrs_name[new_name] = self
}

func (self *Player) SetData(data *PlayerData) {
	self.sid = uint32(query_avail_slot_index())
	self.data = data

	_plrs_sid[self.sid] = self
	_plrs_pid[data.Pid] = self
	_plrs_name[data.Name] = self

	self.on_after_load()
}

// 后期优化:保存index，每次从index处查找
func query_avail_slot_index() int {
	for i := _index; i < MAX_PLAYER_COUNT; i++ {
		if _plrs_sid[i] == nil {
			_index++
			if _index >= MAX_PLAYER_COUNT {
				_index = 1
			}
			return i
		}
	}
	log.Error("query_avail_slot_index FAILED")
	return 0
}

func GetPlayerBySid(sid int) *Player {
	if sid > 0 && sid < MAX_PLAYER_COUNT {
		return _plrs_sid[sid]
	}
	return nil
}

func GetPlayerByPid(pid string) *Player {
	plr, ok := _plrs_pid[pid]
	if !ok {
		return nil
	}
	return plr
}

func GetPlayerByName(name string) *Player {
	plr, ok := _plrs_name[name]
	if !ok {
		return nil
	}
	return plr
}

func IsOnLine(pid string) bool {
	if plr, ok := _plrs_online[pid]; ok {
		return plr != nil
	}
	return false
}

func EachPlayer(fn func(*Player)) {
	for _, plr := range _plrs_online {
		if plr != nil {
			fn(plr)
		}
	}
}

// ============================================================================

func LoadData() {
	// 加载DB中所有的玩家
	arr := []*PlayerData{}
	err := db_mgr.GetDB().GetAllObjects(
		db_mgr.Table_name_player,
		&arr,
	)
	if err != nil && err != mgo.ErrNotFound {
		log.Error("load all PlayerData err :", err)
		return
	}

	for _, data := range arr {
		plr := NewPlayer()
		plr.SetData(data)
	}

	log.Info("[%d] player loaded !", len(arr))
}

func SaveData() {
	// 所有玩家存盘
	// TODO
}

// ============================================================================

func init(){
	event.On(constant.Evt_Auth, func (id uint32, args ...interface{})  {
	// 检测玩家是否已登录

	// acct string,
	// s ISession

	if IsLogin(acct) {
		return false
	}

	// 内存中查找玩家
	plr := GetPlayerByAcct(acct)
	if plr == nil {
		// 新建玩家
		plr = NewPlayer()
		data := CreatePlayerData(acct)
		plr.SetData(data)
		plr.Save()
	}

	plr.Init()

	s.SetPlayer(plr)
	plr.SetSession(s)
})



