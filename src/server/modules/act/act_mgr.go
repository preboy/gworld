package act

import (
	"strings"
	"time"

	"core/log"
	"core/utils"

	"server/app"
	"server/config"
)

// ---------------------------------------------------------

type IAct interface {
	// impl by ActBase
	add_term(*act_term_t)
	check_term() bool

	// impl by real Act
}

var (
	_acts map[int]IAct
)

// ------------------------------------------------------------------------------------
// impl for IAct & Base for real act

type ActBase struct {
	Id    int
	Data  interface{}
	Stage int32 // 0:当前关闭 1:当前打开
	Key   int64 // 如果开始时间(OpenSec)未变，则表示活动仍在同一期

	terms []*act_term_t
}

type act_term_t struct {
	Seq      int
	OpenSec  int64 // 开启时间(单位：秒)
	CloseSec int64 // 结束时间
}

// ------------------------------------------------------------------------------------

func Open() {

	parse_act_config()

	load_act_data()
}

func Close() {

}

// ------------------------------------------------------------------------------------

func RegAct(aid int, act IAct) {
	if _, ok := _acts[aid]; ok {
		log.Warning("activity repeated register, aid =", aid)
		return
	}

	_acts[aid] = act
}

func FindAct(aid int32) IAct {
	return nil
}

func IsOpen(aid int32) bool {
	return false
}

func EachAct(f func(IAct)) {
	for _, act := range _acts {
		f(act)
	}
}

// ------------------------------------------------------------------------------------

func parse_config_date(date string) int64 {
	date = strings.Trim(date, " ")

	if strings.HasPrefix(date, "@") {
		t := time.Unix(app.GetServerData().ServerOpenTime, 0)
		return utils.ParseRelativeTime(t, date).Unix()
	} else {
		return utils.ParseTime(date).Unix()
	}
}

// 加载配置，解析日期
func parse_act_config() {
	config.GetActivityConf().ForEach(func(item *config.ActivityItem) {
		act := _acts[item.Id]
		if act == nil {
			log.Warning("NOT IMPL activity:", item.Id, item.Name)
			return
		}

		act.add_term(&act_term_t{
			Seq:      item.Seq,
			OpenSec:  parse_config_date(item.Open),
			CloseSec: parse_config_date(item.Close),
		})
	})

	// period checking
	overlap := false
	for _, act := range _acts {
		if !act.check_term() {
			overlap = true
		}
	}

	if overlap {
		panic("activity: parse_act_config")
	}
}

func load_act_data() {

}

func save_act_data() {

}
