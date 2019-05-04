package act

import (
	"time"

	"core/log"
)

// ============================================================================

type act_term_t struct {
	Seq      int32
	OpenSec  int64 // 开启时间(单位：秒)
	CloseSec int64 // 结束时间
}

type ActBase struct {
	Id      int32                  `bson:"Id"`
	Status  int32                  `bson:"Status"` // 0:当前关闭 1:当前打开
	Key     int64                  `bson:"Key"`    // 如果开始时间(OpenSec)未变，则表示活动仍在同一期
	DataSvr interface{}            `bson:"DataSvr"`
	DataPlr map[string]interface{} `bson:"DataPlr"`

	terms []*act_term_t `bson:"-"`
}

// ============================================================================
// impl for IAct
// Base for real act

func (self *ActBase) get_id() int32 {
	return self.Id
}

func (self *ActBase) set_id(id int32) {
	self.Id = id
}

func (self *ActBase) get_key() int64 {
	return self.Key
}

func (self *ActBase) set_key(key int64) {
	self.Key = key
}

func (self *ActBase) get_status() int32 {
	return self.Status
}

func (self *ActBase) set_status(status int32) {
	self.Status = status
}

// ============================================================================

func (self *ActBase) set_open(key int64) {
	a := FindAct(self.Id)

	self.DataSvr = a.NewSvrData()
	self.DataPlr = make(map[string]interface{})

	self.Status = 1
	self.Key = key

	a.OnOpen()
}

func (self *ActBase) set_close() {
	a := FindAct(self.Id)

	a.OnClose()

	self.Key = 0
	self.Status = 0
}

func (self *ActBase) IsOpen() bool {
	return self.Status == 1
}

// ============================================================================

func (self *ActBase) add_term(term *act_term_t) {
	self.terms = append(self.terms, term)
}

func (self *ActBase) check_terms() bool {
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

func (self *ActBase) get_key_term() int64 {
	sec := time.Now().Unix()

	for _, term := range self.terms {
		if sec >= term.OpenSec && sec < term.CloseSec {
			return term.OpenSec
		}
	}

	return 0
}

// ============================================================================

func (self *ActBase) GetSvrDataRaw() interface{} {
	if self.DataSvr == nil {
		a := FindAct(self.Id)
		self.DataSvr = a.NewSvrData()
	}

	return self.DataSvr
}

func (self *ActBase) SetSvrDataRaw(data interface{}) {
	if data != nil {
		self.DataSvr = data
	}
}

func (self *ActBase) GetPlrDataRaw() map[string]interface{} {
	if self.DataPlr == nil {
		self.DataPlr = make(map[string]interface{})
	}

	return self.DataPlr
}

func (self *ActBase) SetPlrDataRaw(data map[string]interface{}) {
	if data == nil {
		self.DataPlr = data
	}
}

func (self *ActBase) GetPersonalDataRaw(pid string) interface{} {
	d, ok := self.GetPlrDataRaw()[pid]
	if !ok {
		a := FindAct(self.Id)
		d = a.NewPlrData()
		self.DataPlr[pid] = d
	}

	return d
}

// ============================================================================

func (self *ActBase) NewSvrData() interface{} {
	log.Error("ACT(%d) NOT IMPL `NewSvrData`", self.Id)
	return nil
}

func (self *ActBase) NewPlrData() interface{} {
	log.Error("ACT(%d) NOT IMPL `NewPlrData`", self.Id)
	return nil
}

// ============================================================================
// holdplace

func (self *ActBase) OnOpen() {
}

func (self *ActBase) OnClose() {
}
