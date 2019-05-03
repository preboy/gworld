package act_01

import (
	"game/app"
	"game/constant"
	"game/modules/act"
)

type act_t struct {
	act.ActBase
}

type data_svr_t struct {
}

type data_plr_t struct {
}

var _this_act = &act_t{}

// ============================================================================

func init() {
	println("fucked")
	act.RegAct(constant.ActId_1, _this_act)
}

func (self *act_t) NewSvrData() interface{} {
	return &data_svr_t{}
}

func (self *act_t) NewPlrData() interface{} {
	return &data_plr_t{}
}

// ============================================================================
// DO NOT EDIT

func (self *act_t) GetSvrData() *data_svr_t {
	return self.GetSvrDataRaw().(*data_svr_t)
}

func (self *act_t) GetPlrData(plr app.IPlayer) *data_plr_t {
	return self.GetPersonalDataRaw(plr.GetId()).(*data_plr_t)
}

// ============================================================================
