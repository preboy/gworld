package act_01

import (
	"server/app"
	"server/constant"
	"server/modules/act"
)

type act_t struct {
	act.ActBase
}

type data_svr_t struct {
}

type data_plr_t struct {
}

var act_01 = &act_t{}

// ------------------------------------------------------------------------------------

func init() {
	act.RegAct(constant.ActId_1, act_01)
}

func (self *act_t) NewSvrData() interface{} {
	return &data_svr_t{}
}

func (self *act_t) NewPlrData() interface{} {
	return &data_svr_t{}
}

// ------------------------------------------------------------------------------------
// DO NOT EDIT

func (self *act_t) GetSvrData() *data_svr_t {
	return self.GetRawSvrData().(*data_svr_t)
}

func (self *act_t) GetPlrData(plr app.IPlayer) *data_plr_t {
	return self.GetRawPlrData(plr.GetId()).(*data_plr_t)
}

// ------------------------------------------------------------------------------------
