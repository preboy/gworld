package lobby

import (
	"gworld/ddz/loop"
)

var (
	_matches = map[uint32]*Match{}
	_pids    = []string{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		update()

		for k, v := range _matches {
			if v.Over() {
				delete(_matches, k)
				break
			}

			v.OnUpdate()
		}
	})
}

// ----------------------------------------------------------------------------
// local

func update() {
	create_match()
}

// 人数够了就创建一场斗地主比赛
func create_match() {
	if len(_pids) >= 3 {
		m := NewMatch()
		m.Init(_pids[:3])

		_pids = _pids[3:]
		_matches[m.ID] = m
	}
}

// ----------------------------------------------------------------------------
// export

func Queue(pid string) bool {
	// in queue
	for _, v := range _pids {
		if v == pid {
			return false
		}
	}

	// in match
	for _, v := range _matches {
		if v.Exist(pid) {
			return false
		}
	}

	return true
}

func Init() {
}

func Release() {
}
