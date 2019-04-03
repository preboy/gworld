package config

import (
	"core/log"
)

// ============================================================================

type DropItem struct {
	Prob uint32 `json:"prob"`
	Id   uint32 `json:"id"`
	Cnt  uint64 `json:"cnt"`
}

type Drop struct {
	DropId uint32      `json:"dropId"`
	Prob   []*DropItem `json:"prob"`
	Weight []*DropItem `json:"weight"`
	Cond   []*DropItem `json:"cond"`
	CondId uint32      `json:"condId"`
	WTotal uint32      // Weight项的权重之后
}

type DropTable struct {
	items map[uint32]*Drop
}

// ============================================================================

var (
	DropConf = &DropTable{}
)

// ============================================================================

func (self *DropTable) Load() bool {
	file := "Drop.json"
	var arr []*Drop

	if !load_json_as_arr(C_Config_Path+file, &arr) {
		return false
	}

	self.items = make(map[uint32]*Drop)
	for _, v := range arr {
		key := v.DropId
		self.items[key] = v
	}

	self.drop_calc_total_weight()

	log.Info("load [ %s ] OK", file)
	return true
}

func (self *DropTable) Query(dropid uint32) *Drop {
	return self.items[dropid]
}

func (self *DropTable) Items() map[uint32]*Drop {
	return self.items
}

// ============================================================================

func (self *DropTable) drop_calc_total_weight() {
	for _, item := range self.items {
		for _, v := range item.Weight {
			item.WTotal += v.Prob
		}
	}
}
