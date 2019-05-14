package act_01

import (
	"core/event"
	"game/constant"
	"game/modules/act"
)

type act_t struct {
	act.ActBase
}

var (
	_this_act = &act_t{}
)

// ============================================================================

type data_svr_t struct {
	Total int32
}

type data_plr_t struct {
	LoginCnt int32
}

// ============================================================================

func init() {
	act.RegAct(constant.ActId_1, _this_act)

	event.On(constant.Evt_Plr_Login, func(evt uint32, args ...interface{}) {
		if !_this_act.IsOpen() {
			return
		}

		pid := args[0].(string)

		svr_data := _this_act.GetSvrData()
		plr_data := _this_act.GetPlrData(pid)

		svr_data.Total++
		plr_data.LoginCnt++
	})
}

// ============================================================================
// DO NOT EDIT

func (self *act_t) NewSvrData() interface{} {
	return &data_svr_t{}
}

func (self *act_t) NewPlrData() interface{} {
	return &data_plr_t{}
}

func (self *act_t) GetSvrData() *data_svr_t {
	return self.GetSvrDataRaw().(*data_svr_t)
}

func (self *act_t) GetPlrData(pid string) *data_plr_t {
	return self.GetPlrDataRaw(pid).(*data_plr_t)
}

// ============================================================================
