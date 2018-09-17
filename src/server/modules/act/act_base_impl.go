package act

import (
	"core/log"
)

// ------------------------------------------------------------------------------------
// impl for IAct
// Base for real act

func (self *ActBase) add_term(term *act_term_t) {
	self.terms = append(self.terms, term)
}

func (self *ActBase) check_term() bool {
	pass := true

	l := len(self.terms)
	for i := 0; i < l; i++ {
		for j := i + 1; i < l; j++ {
			if (self.terms[i].OpenSec >= self.terms[j].OpenSec && self.terms[i].OpenSec < self.terms[j].CloseSec) ||
				(self.terms[j].OpenSec >= self.terms[i].OpenSec && self.terms[j].OpenSec < self.terms[i].CloseSec) {
				log.Warning("活动开放时间有重叠:", self.Id, self.terms[i].Seq, self.terms[j].Seq)
				pass = false
			}
		}
	}

	return pass
}
