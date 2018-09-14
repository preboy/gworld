package act

import (
	"core/log"
	"core/utils"

	"server/game"
	"server/game/config"
)

// ---------------------------------------------------------

type IActBase interface {
	GetId() int
}

type act_term_t struct {
	Seq      uint32
	OpenSec  int64 // 开启时间(单位：秒)
	CloseSec int64 // 结束时间
}

type act_t struct {
	Id    uint32
	Data  interface{}
	Stage int32 // 0:当前关闭 1:当前打开
	Key   int64 // 如果开始时间(OpenSec)未变，则表示活动仍在同一期

	obj   IActBase
	terms []*act_term_t
}

var (
	_acts map[int]*act_t
)

// ------------------------------------------------------------------------------------

func RegisterActivity(aid int, act IActBase) {
	if _, ok := _acts[aid]; ok {
		log.Warning("activity repeated register, aid =", aid)
		return
	}

	_acts[aid] = &act_t{
		Id:  aid,
		Obj: act,
	}
}

func Open() {

	parse_act_config()

	load_act_data()
}

func Close() {

}

// ------------------------------------------------------------------------------------

func parse_config_date(date string) int {
	date = strings.Trim(date, " ")

	if strings.HasPrefix(date, "@") {
		t = time.Unix(game.GetServerData().ServerOpenTime, 0)
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

		term := &act_term_t{
			Seq:      item.Seq,
			OpenSec:  parse_config_date(item.Open),
			CloseSec: parse_config_date(item.Close),
		}

		act.terms = append(act.terms, term)
	})

	// period checking
	overlap := false
	for _, act := range _acts {
		l := len(act.terms)
		for i := 0; i < l; i++ {
			for j := i + 1; i < l; j++ {
				if (act.terms[i].OpenSec >= act.terms[j].OpenSec && act.terms[i].OpenSec < act.terms[j].CloseSec) ||
					(act.terms[j].OpenSec >= act.terms[i].OpenSec && act.terms[j].OpenSec < act.terms[i].CloseSec) {
					log.Warning("活动开放时间有重叠", act.Id, act.terms[i].Seq, act.terms[j].Seq)
					overlap = true
				}
			}
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
