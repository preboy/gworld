package microsvr

import (
	"sync"
)

var (
	svrs = map[uint16]*Svr{}
)

func FindSvrByOpcode(opcode uint16) *Svr {
	if opcode != 0 {
		return svrs[opcode/100]
	}
	return nil
}

func FindSvrByName(name string) *Svr {
	for _, v := range svrs {
		if v.mod.GetName() == name {
			return v
		}
	}
	return nil
}

func Register(mod IMod) *Svr {
	svr := &Svr{
		w:    &sync.WaitGroup{},
		mod:  mod,
		quit: make(chan bool),
		msgc: make(chan *msg, 0x1000),
	}

	key, _ := mod.GetOpcodeRange()
	svrs[key] = svr

	return svr
}
