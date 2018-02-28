package player

const (
	MAX_PLAYER_COUNT = 0X4000
)

var (
	_index int
)

var (
	// 在内存中的玩家，包括主动上线、从DB中被拉上线的
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

func EnterGame(acct string, s ISession) bool {
	// 在内存中查找玩家
	// 在DB中查找玩家
	var plr *Player = CreatePlayer(acct)
	if plr == nil {
		return false
	}
	s.SetPlayer(plr)
	plr.SetSession(s)
	plr.Go()
	return true
}

// ------------- local function -------------
func CreatePlayer(acct string) *Player {
	plr := GetPlayerByAcct(acct)
	if plr == nil {
		plr = NewPlayer()

		ok, data := LoadPlayerData(acct)
		if !ok {
			data = CreatePlayerData(acct)
		}

		// 新的对象入坑
		plr.sid = uint32(query_avail_slot_index())
		plr.data = data

		_plrs_sid[plr.sid] = plr
		_plrs_pid[data.Pid] = plr
		_plrs_name[data.Name] = plr
		_plrs_acct[data.Acct] = plr

		if ok {
			plr.on_after_load()
		}
	}

	return plr
}

func EachPlayer(fn func(*Player)) {
	for _, plr := range _plrs_acct {
		fn(plr)
	}
}
