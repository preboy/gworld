package player

import (
	"fmt"

	"core/event"
	"core/log"
	"game/app"
	"game/constant"
	"game/dbmgr"
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

// ============================================================================

func init() {
	event.On(constant.Evt_Auth, func(id uint32, args ...interface{}) {
		// 检测玩家是否已登录

		pid := args[0].(string)
		acct := args[1].(string)
		sess := args[2].(ISession)
		first := false

		plr := GetPlayer(pid)
		if plr == nil {
			plr = NewPlayer()

			data := LoadPlayerDataFromDB(pid)
			if data == nil {
				data = CreatePlayerData(acct)
				first = true
			}

			plr.SetData(data)
			plr.Init()
		} else {
			if plr.IsOnLine() {
				plr.Logout()
			}
		}

		plr.SetSession(sess)
		plr.Login(first)
	})
}

func (self *Player) SetName(name string) {
	// TODO: checking valid & repeat

	old_name := self.data.Name
	new_name := fmt.Sprintf("%s-%s", app.GetGameId(), name)

	self.data.Name = new_name

	_plrs_name[old_name] = nil
	_plrs_name[new_name] = self
}

func (self *Player) SetData(data *PlayerData) {
	sid := uint32(query_avail_slot_index())
	if sid == 0 {
		return
	}

	self.sid = sid
	self.data = data

	_plrs_sid[self.sid] = self
	_plrs_pid[data.Pid] = self
	_plrs_name[data.Name] = self
}

// ============================================================================

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

func GetPlayer(pid string) *Player {
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

func EachOnlinePlayer(fn func(*Player)) {
	for _, plr := range _plrs_online {
		if plr != nil {
			fn(plr)
		}
	}
}

// ============================================================================

func LoadData() {
	// 加载DB中某些的玩家
	arr := []*PlayerData{}
	err := dbmgr.GetDB().GetAllObjects(
		dbmgr.Table_name_player,
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
