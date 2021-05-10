package lobby

import (
	"gworld/ddz/comp"
	"gworld/ddz/loop"
)

var (
	_matches = map[uint32]comp.IMatch{}
	_pids    = []string{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {

		for k, m := range _matches {
			if m.IsOver() {
				delete(_matches, k)
				break
			}

			m.OnUpdate()
		}
	})
}

// ----------------------------------------------------------------------------
// export

func Init() {
}

func Release() {
}

func AddMatch(m comp.IMatch) {
	_matches[m.GetMID()] = m
}

func DelMatch(mid uint32) {
	delete(_matches, mid)
}

func GetMatch(mid uint32) comp.IMatch {
	return _matches[mid]
}

func GetMatchByName(name string) comp.IMatch {
	for _, m := range _matches {
		if m.GetName() == name {
			return m
		}
	}

	return nil
}

func OnMessage(pid string, req comp.IMessage, res comp.IMessage) {
}
