package player

import (
	"fmt"
	"time"
)

// 处理离线时间段的搁置操作
func (self *Player) pursue() {

	now := time.Now()
	lst := self.data.LogoutTs

	if now.Hour() != lst.Hour() {
		self.on_new_hour()
	}
	if now.Day() != lst.Day() {
		self.on_new_day()
	}

	_, week_now := now.ISOWeek()
	_, week_lst := lst.ISOWeek()

	if week_now != week_lst {
		self.on_new_week()
	}
	if now.Month() != lst.Month() {
		self.on_new_month()
	}
	if now.Year() != lst.Year() {
		self.on_new_year()
	}
}
