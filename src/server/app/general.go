package app

import (
	"strconv"
	"sync"
)

var (
	_lock *sync.Mutex
)

func init() {
	_lock = &sync.Mutex{}
}

// 产生一个玩家ID
func GeneralPlayerID() string {
	_lock.Lock()
	defer func() {
		_lock.Unlock()
	}()

	sd := GetServerData()
	sc := GetServerConfig()

	sd.IdSeq++

	return strconv.FormatUint(uint64(sc.Server_id), 10) + strconv.FormatUint(uint64(sd.IdSeq), 10)
}

func GeneralPlayerName(pid string) string {
	return "lord_" + pid
}
