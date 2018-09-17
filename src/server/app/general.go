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
func GeneralPlayerID() uint64 {
	_lock.Lock()
	defer func() {
		_lock.Unlock()
	}()

	sd := GetServerData()
	sc := GetServerConfig()

	sd.IdSeq++
	// 64 bit:  serverid(16)+_(16)+idseq(32)
	return uint64(sc.Server_id<<48) + uint64(sd.IdSeq)

}

func GeneralPlayerName(pid uint64) string {
	return "lord_" + strconv.FormatUint(pid, 10)
}
