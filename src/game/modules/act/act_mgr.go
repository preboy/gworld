package act

import (
	"strings"

	"core/db"
	"core/log"
	"core/utils"
	"game/app"
	"game/config"
	"game/dbmgr"
	"gopkg.in/mgo.v2/bson"
)

// ============================================================================

var (
	_acts = make(map[int32]IAct, 128)
)

// ============================================================================

type IAct interface {
	get_id() int32
	set_id(id int32)

	get_key() int64
	set_key(key int64)

	get_status() int32
	set_status(status int32)

	set_open(key int64)
	set_close()
	IsOpen() bool

	add_term(*act_term_t)
	check_terms() bool
	get_key_term() int64

	GetActBase() *ActBase

	GetSvrDataRaw() interface{}
	SetSvrDataRaw(data interface{})

	GetTabDataRaw() map[string]interface{}
	SetTabDataRaw(data map[string]interface{})

	GetPlrDataRaw(pid string) interface{}

	NewPlrData() interface{}
	NewSvrData() interface{}

	OnOpen()
	OnClose()
}

// ============================================================================

func Open() {
	parse_act_config()
	load_act_data()
	check_act_status()
}

func Close() {
	save_act_data()
}

// ============================================================================

func RegAct(id int32, a IAct) {
	if _, ok := _acts[id]; ok {
		log.Warning("ACT(%d) repeated register !!!", id)
		return
	}

	a.set_id(id)

	_acts[id] = a
}

func FindAct(id int32) IAct {
	return _acts[id]
}

func ActIsOpen(id int32) bool {
	if a, ok := _acts[id]; ok {
		return a.IsOpen()
	}

	return false
}

func EachAct(f func(IAct)) {
	for _, a := range _acts {
		f(a)
	}
}

// ============================================================================

func parse_config_date(date string) int64 {
	date = strings.Trim(date, " ")

	if strings.HasPrefix(date, "@") {
		return utils.ParseRelativeTime(app.GetServerData().ServerOpenTime, date).Unix()
	} else {
		return utils.ParseTime(date).Unix()
	}
}

// 加载配置，解析日期
func parse_act_config() {
	config.ActivityConf.ForEach(func(item *config.Activity) {
		a := _acts[item.Id]
		if a == nil {
			log.Warning("NOT IMPL activity: {id=%v, name=%v}", item.Id, item.Name)
			return
		}

		a.add_term(&act_term_t{
			Seq:      item.Seq,
			OpenSec:  parse_config_date(item.Open),
			CloseSec: parse_config_date(item.Close),
		})
	})

	// period checking
	for _, a := range _acts {
		a.check_terms()
	}
}

// ============================================================================

func convert_svr_data(a IAct, in interface{}) (out interface{}) {
	out = a.NewSvrData()

	if in == nil {
		return
	}

	data, err := bson.Marshal(in)
	if err != nil {
		log.Error("bson.Marshal err = %v", err)
		return
	}

	bson.Unmarshal(data, out)

	return
}

func convert_tab_data(a IAct, in map[string]interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})

	if in == nil {
		return
	}

	for k, v := range in {

		d := a.NewPlrData()

		data, err := bson.Marshal(v)
		if err != nil {
			log.Error("bson.Marshal err = %v", err)
			return
		}

		bson.Unmarshal(data, d)

		out[k] = d
	}

	return
}

// ============================================================================

func load_act_data() {

	type act_t struct {
		_id     int32    `bson:"_id"`
		ActBase *ActBase `bson:"actbase"`
	}

	var arr []*act_t

	err := dbmgr.GetDB().GetAllObjects(dbmgr.Table_name_activity, &arr)
	if err != nil {
		if db.IsNotFound(err) {
			log.Info("Loading < %v >, IsNotFound !", dbmgr.Table_name_activity)
		} else {
			log.Fatal("Loading < %v >  Fatal !!!", dbmgr.Table_name_activity)
			return
		}
	} else {

		for _, v := range arr {

			if v.ActBase == nil {
				continue
			}

			if a, ok := _acts[v.ActBase.Id]; ok {
				a.set_key(v.ActBase.Key)
				a.set_status(v.ActBase.Status)
				a.SetSvrDataRaw(convert_svr_data(a, v.ActBase.DataSvr))
				a.SetTabDataRaw(convert_tab_data(a, v.ActBase.DataTab))
			}
		}
	}
}

func save_act_data() {
	for _, a := range _acts {
		dbmgr.GetDB().Upsert(dbmgr.Table_name_activity, a.get_id(), a)
	}
}

func check_act_status() {
	for _, a := range _acts {
		key := a.get_key_term()

		if !a.IsOpen() {
			if key == 0 {
				// also closed, do nothing
			} else {
				// new team, set open
				a.set_open(key)
			}
		} else {
			if key == 0 {
				// closed
				a.set_close()
			} else {
				if key == a.get_key() {
					// some team
				} else {
					// another team
					a.set_close()
					a.set_open(key)
				}
			}
		}
	}
}

func OnLoopUpdate() {
	// schedule
	check_act_status()
}
