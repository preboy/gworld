package act_01

import (
	"server/constant"
	"server/modules/act"
)

type act_t struct {
	act.ActBase
}

var act_01 = &act_t{}

// ------------------------------------------------------------------------------------

func init() {
	act.RegAct(constant.ActId_1, act_01)
}
