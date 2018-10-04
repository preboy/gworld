package player

import (
	"core/log"
	"server/app"
)

const (
	MAX_PLAYER_COUNT = 0X4000
)

var (
	_index int = 1
)

/*
服务器启动时从DB中加载所有的玩家到内存，存入_plrs_pid、_plrs_name、_plrs_acct
新玩家立即存盘，并加载到内存
登录期间的玩家存入_plrs_live、_plrs_sid，并设置玩家的run状态，玩家下线回收上述参数
玩家上线通过_plrs_acct登录是否已登录
游戏中可通过_plrs_live、_plrs_sid快速查找在线玩家
游戏中可通过_plrs_acct、_plrs_name、_plrs_pid查找所有玩家
*/
var (
	// 在内存中的玩家，包括主动上线、从DB中被拉上线的
	_plrs_sid  [MAX_PLAYER_COUNT]*Player // 运行中的玩家
	_plrs_pid  map[string]*Player        // pid
	_plrs_name map[string]*Player        // name
	_plrs_acct map[string]*Player        // acct
	_plrs_live map[string]*Player        // 已登录的玩家
)

func init() {
	_plrs_sid = [MAX_PLAYER_COUNT]*Player{}
	_plrs_pid = make(map[string]*Player, MAX_PLAYER_COUNT)
	_plrs_name = make(map[string]*Player, MAX_PLAYER_COUNT)
	_plrs_acct = make(map[string]*Player, MAX_PLAYER_COUNT)
	_plrs_live = make(map[string]*Player, MAX_PLAYER_COUNT)
}

func (self *Player) SetName(name string) {
	old_name := self.data.Name
	self.data.SetName(name)
	new_name := self.data.Name
	_plrs_name[old_name] = nil
	_plrs_name[new_name] = self
}

func (self *Player) AssociateData(data *PlayerData) {
	self.sid = uint32(query_avail_slot_index())
	self.data = data

	_plrs_sid[self.sid] = self
	_plrs_pid[data.Pid] = self
	_plrs_name[data.Name] = self
	_plrs_acct[data.Acct] = self

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

func GetPlayerByAcct(acct string) *Player {
	plr, ok := _plrs_acct[acct]
	if !ok {
		return nil
	}
	return plr
}

func IsLogin(acct string) bool {
	plr, ok := _plrs_live[acct]
	if ok {
		return plr != nil
	}
	return false
}

func EnterGame(acct string, s ISession) bool {
	// 检测玩家是否已登录
	if IsLogin(acct) {
		return false
	}

	// 内存中查找玩家
	plr := GetPlayerByAcct(acct)
	if plr == nil {
		// 新建玩家
		plr = NewPlayer()
		data := CreatePlayerData(acct)
		plr.AssociateData(data)
		plr.Save()
	}

	plr.Init()

	plr.server_id = app.GetServerConfig().Server_id

	s.SetPlayer(plr)
	plr.SetSession(s)
	plr.Go()

	return true
}

func EachPlayer(f func(*Player)) {
	for _, plr := range _plrs_live {
		if plr != nil {
			f(plr)
		}
	}
}
