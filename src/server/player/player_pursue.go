package player

import (
	"time"
)

// 处理离线时间段的搁置操作
func (self *Player) pursue() {
	now := time.Now().UnixNano() / (1000 * 1000)
	if now >= self.last_update+100 {

	}

	dt := (now - self.last_update) / 1000
	self.on_pursue(dt)

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

func (self *Player) on_pursue(sec int64) {

}

func (self *Player) on_new_hour() {

}

func (self *Player) on_new_day() {

}

func (self *Player) on_new_week() {

}

func (self *Player) on_new_month() {

}

func (self *Player) on_new_year() {

}
