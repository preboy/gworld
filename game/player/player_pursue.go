package player

import (
	"time"

	"gworld/core/event"
	"gworld/game/constant"
)

// 处理离线时间段的搁置操作
func (self *Player) pursue() {
	now := time.Now()
	lst := self.data.LogoutTs

	if now.Hour() != lst.Hour() {
		self.on_span_hour()
	}
	if now.Day() != lst.Day() {
		self.on_span_day()
	}

	_, week_now := now.ISOWeek()
	_, week_lst := lst.ISOWeek()

	if week_now != week_lst {
		self.on_span_week()
	}
	if now.Month() != lst.Month() {
		self.on_span_month()
	}
	if now.Year() != lst.Year() {
		self.on_span_year()
	}
}

// ============================================================================

func (self *Player) on_span_hour() {
}

func (self *Player) on_span_day() {
	event.Fire(constant.Evt_Plr_DataReset, self)
}

func (self *Player) on_span_week() {
}

func (self *Player) on_span_month() {
}

func (self *Player) on_span_year() {
}
