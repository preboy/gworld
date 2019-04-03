package act

import (
	"time"

	"core/log"
	"core/utils"
)

// ============================================================================
// impl for IAct
// Base for real act

func (self *ActBase) add_term(term *act_term_t) {
	self.terms = append(self.terms, term)
}

func (self *ActBase) check_term() bool {
	pass := true

	l := len(self.terms)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if (self.terms[i].OpenSec >= self.terms[j].OpenSec && self.terms[i].OpenSec < self.terms[j].CloseSec) ||
				(self.terms[j].OpenSec >= self.terms[i].OpenSec && self.terms[j].OpenSec < self.terms[i].CloseSec) {
				log.Warning("活动开放时间有重叠:", self.Id, self.terms[i].Seq, self.terms[j].Seq)
				pass = false
			}
		}
	}

	return pass
}

func (self *ActBase) GetRawSvrData() interface{} {
	if self.DataSvr == nil {
		self.DataSvr = self.NewSvrData()
	}
	return self.DataSvr
}

func (self *ActBase) GetRawPlrData() map[string]interface{} {
	if self.DataPlr == nil {
		self.DataPlr = make(map[string]interface{})
	}
	return self.DataPlr
}

func (self *ActBase) GetRawPlrTable(id string) interface{} {
	d, ok := self.DataPlr[id]
	if !ok {
		d = self.NewPlrData(id)
		self.DataPlr[id] = d
	}
	return d
}

func (self *ActBase) get_id() int {
	return self.Id
}

func (self *ActBase) get_stage() int32 {
	return self.Stage
}

func (self *ActBase) get_key() int64 {
	return self.Key
}

func (self *ActBase) is_open() bool {
	return self.Stage == 1
}

func (self *ActBase) get_key_curr() int64 {
	sec := time.Now().Unix()
	for _, term := range self.terms {
		if sec >= term.OpenSec && sec < term.CloseSec {
			return term.OpenSec
		}
	}
	return 0
}

func (self *ActBase) set_close() {
	utils.ExecuteSafely(func() {
		self.OnClose()
	})

	self.Stage = 0
	self.Key = 0
}

func (self *ActBase) set_open(key int64) {
	self.DataSvr = self.NewSvrData()
	self.DataPlr = make(map[string]interface{})

	self.Stage = 1
	self.Key = key

	self.OnOpen()
}

// ============================================================================

func (self *ActBase) NewSvrData() interface{} {
	return nil
}

func (self *ActBase) NewPlrData(id string) interface{} {
	return nil
}

// ============================================================================
// holdplace

func (self *ActBase) OnOpen() {

}
func (self *ActBase) OnClose() {

}
