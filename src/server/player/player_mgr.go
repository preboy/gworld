package player

import (
	"math/rand"
)

const (
	MAX_PLAYER_COUNT = 0X4000
)

var (
	_index int
)

var (
	_plrs_sid  = [MAX_PLAYER_COUNT]*Player{}
	_plrs_pid  = make(map[uint64]*Player)
	_plrs_name = make(map[string]*Player)
	_plrs_acct = make(map[string]*Player)
)

// 后期优化:保存index，每次从index处查找
func query_avail_slot_index() int {
	for i := 0; i < MAX_PLAYER_COUNT; i++ {
		if _plrs_sid[i] == nil {
			return i
		}
	}
	return -1
}

func GetPlayerBySid(sid int) *Player {
	if sid >= 0 && sid < MAX_PLAYER_COUNT {
		return _plrs_sid[sid]
	}
	return nil
}

func GetPlayerByPid(pid uint64) *Player {
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

func EnterGame(acct string, s ISession) {
	// 在内存中查找玩家
	// 在DB中查找玩家
	var plr *Player = nil
	plr = GetPlayerByAcct(acct)
	if plr == nil {
		plr = CreatePlayer()
	}

	s.SetPlayer(plr)
	plr.SetSession(s)

}

// ------------- local function -------------
func CreatePlayer() *Player {

	plr := NewPlayer()

	plr.pid = rand.Uint64()
	plr.sid = uint32(query_avail_slot_index())
	plr.name = ""
	plr.acct = ""

	_plrs_sid[plr.sid] = plr
	_plrs_pid[plr.pid] = plr

	_plrs_name[plr.name] = plr
	_plrs_acct[plr.acct] = plr

	return plr
}
