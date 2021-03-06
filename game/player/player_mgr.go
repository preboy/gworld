package player

import (
	"fmt"

	"gworld/core/event"
	"gworld/core/log"
	"gworld/core/work"
	"gworld/game/constant"
	"gworld/game/dbmgr"

	"gopkg.in/mgo.v2"
)

const (
	MAX_PLAYER_COUNT = 0x4000
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
	_plrs_key    = make(map[string]*Player, MAX_PLAYER_COUNT) // key(acct-svr)
	_plrs_name   = make(map[string]*Player, MAX_PLAYER_COUNT) // name
	_plrs_online = make(map[string]*Player, MAX_PLAYER_COUNT) // pid(在线的玩家)
)

// ============================================================================

func init() {
	event.On(constant.Evt_Auth, func(id uint32, args ...interface{}) {
		// 检测玩家是否已登录

		sdk := args[0].(string)
		svr := args[2].(string)
		acct := args[1].(string)
		sess := args[3].(ISession)

		key := fmt.Sprintf("%s-%s", acct, svr)
		if plr, ok := _plrs_key[key]; ok {
			if plr == nil {
				return // loading
			}

			// relogin
			if plr.IsOnLine() {
				plr.Disconnect()
			}

			plr.SetSession(sess)
			plr.Init()
			plr.OnLogin()

		} else {

			_plrs_key[key] = nil

			work.Queue(func() func() {
				data := GetPlayerData(key, acct, svr, sdk)

				return func() {
					plr := NewPlayer().SetData(data)
					plr.SetSession(sess)
					plr.Init()
					plr.OnLogin()
				}
			}, nil)
		}
	})
}

func (self *Player) SetName(name string) {
	// TODO: checking valid & repeat

	old_name := self.data.Name
	new_name := fmt.Sprintf("%s-%s", self.data.Svr, name)

	self.data.Name = new_name

	_plrs_name[old_name] = nil
	_plrs_name[new_name] = self
}

func (self *Player) SetData(data *PlayerData) *Player {
	sid := uint32(query_avail_slot_index())
	if sid == 0 {
		return nil
	}

	self.sid = sid
	self.data = data

	_plrs_sid[self.sid] = self
	_plrs_pid[data.Pid] = self
	_plrs_key[data.Key] = self
	_plrs_name[data.Name] = self

	return self
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
	return _plrs_pid[pid]
}

func GetPlayerByName(name string) *Player {
	return _plrs_name[name]
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
		NewPlayer().SetData(data)
	}

	log.Info("[%d] player loaded !", len(arr))
}

func SaveData() {
	for _, plr := range _plrs_pid {
		plr.Save()
	}
}
