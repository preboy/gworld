package smatch

import (
	"gworld/ddz/lobby/poker"
)

// 一个人本副牌的信息
type deck_data struct {
	cards []poker.Card
}

type gambler struct {
	m *Table

	pid  string
	pos  seat
	data *deck_data

	// stat
	score_total   int // 总分
	win_count     int // 胜次数
	lost_count    int // 败次数
	load_count    int // 地主次数
	peasant_count int // 农民次数
}
