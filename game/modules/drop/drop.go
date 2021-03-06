package drop

import (
	"gworld/core/rand"
	"gworld/game/app"
	"gworld/game/config"
	"gworld/game/modules/cond"
)

const MAX_PROB = 10000

type drop_item struct {
	Id  uint32
	Cnt uint64
}

func Drop(plr app.IPlayer, dropid uint32) (ret []*drop_item) {
	conf := config.DropConf.Query(dropid)
	if conf == nil {
		return
	}

	// 概率掉落
	if len(conf.Prob) > 0 {
		for _, v := range conf.Prob {
			if v.Prob >= MAX_PROB || rand.RandomHitn(int(v.Prob), MAX_PROB) {
				ret = append(ret, &drop_item{
					Id:  v.Id,
					Cnt: v.Cnt,
				})
			}
		}
	}

	// 权重掉落
	if len(conf.Weight) > 0 {
		cnt := uint32(0)
		get := rand.Uint32(int32(conf.WTotal)) + 1

		for _, v := range conf.Weight {
			cnt += v.Prob
			if cnt >= get {
				ret = append(ret, &drop_item{
					Id:  v.Id,
					Cnt: v.Cnt,
				})
				break
			}
		}
	}

	// 条件掉落
	if conf.CondId != 0 && cond.Check(plr, conf.CondId) {
		for _, v := range conf.Cond {
			if v.Prob >= MAX_PROB || rand.RandomHitn(int(v.Prob), MAX_PROB) {
				ret = append(ret, &drop_item{
					Id:  v.Id,
					Cnt: v.Cnt,
				})
			}
		}
	}

	return
}
