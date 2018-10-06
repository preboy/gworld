package act

import (
	"strings"
	"time"

	"core/db"
	"core/log"
	"core/utils"
	"server/app"
	"server/config"
	"server/db_mgr"
)

// ============================================================================

type IAct interface {
	get_id() int
	get_key() int64
	get_stage() int32

	// impl by ActBase
	is_open() int64
	add_term(*act_term_t)
	check_term() bool

	GetRawSvrData() interface{}
	GetRawPlrData() map[string]interface{}
	GetRawPlrTable(id string) interface{}
}

var (
	_acts map[int]IAct
)

func init() {
	_acts = make(map[int]IAct)
}

// ============================================================================
// impl for IAct & Base for real act

type ActBase struct {
	Id      int
	DataSvr interface{}
	DataPlr map[string]interface{}
	Stage   int32 // 0:当前关闭 1:当前打开
	Key     int64 // 如果开始时间(OpenSec)未变，则表示活动仍在同一期

	terms []*act_term_t
}

type act_term_t struct {
	Seq      int
	OpenSec  int64 // 开启时间(单位：秒)
	CloseSec int64 // 结束时间
}

// ============================================================================

func Open() {
	parse_act_config()
	load_act_data()
}

func Close() {
	save_act_data()
}

// ============================================================================

func RegAct(aid int, act IAct) {
	if _, ok := _acts[aid]; ok {
		log.Warning("activity repeated register, aid =", aid)
		return
	}

	_acts[aid] = act
}

func FindAct(id int) IAct {
	return _acts[id]
}

func IsOpen(id int) bool {
	if act, ok := _acts[id]; ok {
		return act.is_open() != 0
	}
	return false
}

func EachAct(f func(IAct)) {
	for _, act := range _acts {
		f(act)
	}
}

// ============================================================================

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
	config.ActivityConf.ForEach(func(item *config.Activity) {
		act := _acts[item.Id]
		if act == nil {
			log.Warning("NOT IMPL activity: {id=%v, name=%v}", item.Id, item.Name)
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

// ============================================================================

func load_act_data() {
	var arr []*ActBase

	err := db_mgr.GetDB().GetAllObjects(db_mgr.Table_name_activitys, &arr)
	if err != nil {
		if db.IsNotFound(err) {
			log.Info("Loading < %v >, IsNotFound !", db_mgr.Table_name_activitys)
		} else {
			log.Fatal("Loading < %v >  Fatal !!!", db_mgr.Table_name_activitys)
			return
		}
	} else {
		for _, v := range arr {
			if act, ok := _acts[v.Id]; ok {
				a := act.(*ActBase)
				a.DataSvr = v.DataSvr
				a.DataPlr = v.DataPlr
				a.Stage = v.Stage
				a.Key = v.Key
			}
		}
	}

	// checking ...
	for _, act := range _acts {
		_ = act
	}

}

func save_act_data() {
	type act_rec_t struct {
		Acts []*ActBase
	}

	rec := &act_rec_t{
		Acts: make([]*ActBase, 0, len(_acts)),
	}

	for _, act := range _acts {
		rec.Acts = append(rec.Acts, &ActBase{
			Id:      act.get_id(),
			DataSvr: act.GetRawSvrData(),
			DataPlr: act.GetRawPlrData(),
			Stage:   act.get_stage(),
			Key:     act.get_key(),
		})
	}

	db_mgr.GetDB().Upsert(db_mgr.Table_name_activitys, 1, rec)
}
