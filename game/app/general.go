package app

import (
	"fmt"
	"sync"
)

var (
	_lock *sync.Mutex
)

func init() {
	_lock = &sync.Mutex{}
}

// 产生一个玩家ID
func GeneralPlayerID() (string, string) {
	_lock.Lock()
	defer func() {
		_lock.Unlock()
	}()

	SeqId := GetServerData().GetSeqId()

	pid := fmt.Sprintf("%s_%05d", GetGameId(), SeqId)
	name := fmt.Sprintf("%s-finder_%05d", GetGameId(), SeqId)

	return pid, name
}
