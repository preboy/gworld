package mod_sample

import (
	"fmt"

	"core/event"
	_ "game/app"
	"game/microsvr"
)

// opcode range:  [n*100, (n*100+99)

const (
	C_Name        = "mod_sample"
	C_OpCodeBegin = 0x1000
	C_OpCodeEnd   = C_OpCodeBegin + 99
)

var (
	svr *microsvr.Svr
	mod = &ModSample{}
)

type ModSample struct {
	sex  int
	name string
}

// ----------------------------------------------------------------------------

func init() {
	svr = microsvr.Register(mod)
}

// ----------------------------------------------------------------------------

func (self *ModSample) GetName() string {
	return C_Name
}

func (self *ModSample) GetOpcodeRange() (uint16, uint16) {
	return C_OpCodeBegin, C_OpCodeEnd
}

func (self *ModSample) OnStart() {
	// load data
}

func (self *ModSample) OnStop() {
	// save data
}

func (self *ModSample) OnUpdate() {
}

func (self *ModSample) OnTimer(id uint64) {
	fmt.Println("Svr.OnTimer:", C_Name, id)
}

func (self *ModSample) OnEvent(evt *event.Event) {
	fmt.Println("Svr.OnEvent:", C_Name, evt)
}
