package player

import (
	"fmt"
	"time"
)

// 处理离线时间段的搁置操作
func (self *Player) pursue() {

	now := time.Now()
	lst := time.Unix(self.last_update/1000, 0)

	self.on_pursue(now.Unix() - self.last_update/1000)

	fmt.Println(now)
	fmt.Println(lst)

	if now.Hour() != lst.Hour() {
		self.on_new_hour()
	}
	if now.Day() != lst.Day() {
		self.on_new_day()
	}
	if now.Weekday() != lst.Weekday() {
		self.on_new_week()
	}
	if now.Month() != lst.Month() {
		self.on_new_month()
	}
	if now.Year() != lst.Year() {
		self.on_new_year()
	}

}

func (self *Player) on_pursue(off_sec int64) {
	println("offline sec:", off_sec, self.data.Name)
}
