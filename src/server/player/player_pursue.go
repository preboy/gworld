package player

import (
	"time"
)

// 处理离线时间段的搁置操作
func (self *Player) pursue() {
	self.on_pursue(time.Now().Unix() - self.last_update/1000)

	end := time.Now()
	bgn := time.Unix(self.last_update/1000, 0)

	if end.Hour() != bgn.Hour() {
		self.on_new_hour()
	}
	if end.Day() != bgn.Day() {
		self.on_new_day()
	}
	if end.Weekday() != bgn.Weekday() {
		self.on_new_week()
	}
	if end.Month() != bgn.Month() {
		self.on_new_month()
	}
	if end.Year() != bgn.Year() {
		self.on_new_year()
	}

}

func (self *Player) on_pursue(off_sec int64) {

}
