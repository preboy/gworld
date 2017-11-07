package player

import (
	_ "fmt"
	"math/rand"
	"net"
	"time"
)

const (
	MAX_PLAYER_COUNT = 0X10000
)

var (
	_plrs_map = make(map[uint32]*Player)
	_plrs_arr = [MAX_PLAYER_COUNT]*Player{}
)

// 后期优化:保存index，每次从index处找
func query_array_index() int {
	for i := 0; i < MAX_PLAYER_COUNT; i++ {
		if !(_plrs_arr[i]) {
			return i
		}
	}
	return -1
}

func CreatePlayer(name string) *Player {

	var plr = NewPlayer()

	plr.pid = rand.Uint32()
	plr.aid = query_array_index()
	plr.name = name
	plr.socket = nil

	_plrs_arr[plr.aid] = plr
	_plrs_map[plr.pid] = plr

	go plr.Loop()

	return plr
}

func GetPlayerByIndex(sid uint32) *Player {
	if _index < MAX_PLAYER_COUNT {
		return _plrs_arr[sid]
	}
	return nil
}

func GetPlayerById(pid uint32) *Player {
	plr, ok := _plrs_map[pid]
	if !ok {
		return nil
	}
	return plr
}
