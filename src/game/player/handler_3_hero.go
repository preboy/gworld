package player

import (
	"core/event"
	"core/rand"
	"core/tcp"
	"game/app"
	"game/config"
	"game/constant"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
)

func handler_HeroLevelupRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.HeroLevelupRequest{}
	res := &msg.HeroLevelupResponse{}
	proto.Unmarshal(packet.Data, req)

	res.ErrorCode = ec.OK

	hero := plr.GetHero(req.Id)

	var lv_old, lv_new uint32

	func() {
		if hero == nil {
			res.ErrorCode = ec.Hero_Not_Activated
			return
		}

		// 不能超过角色等级
		if hero.Lv >= plr.GetLv() {
			res.ErrorCode = ec.Level_Exceed
			return
		}

		// 可否升到下一级
		lvconf := config.LevelupConf.Query(hero.Lv + 1)
		if lvconf == nil {
			res.ErrorCode = ec.Conf_Invalid
			return
		}

		// 道具数量是否足够
		goods := app.NewItemProxy(constant.ItemLog_HeroLvUp).SetArgs(hero.Id, hero.Lv+1)
		for _, v := range lvconf.Needs {
			goods.Sub(v.Id, v.Cnt)
		}
		if !goods.Enough(plr) {
			res.ErrorCode = ec.Item_Not_Enough
			return
		}

		lv_old = hero.Lv
		hero.Lv++
		lv_new = hero.Lv
		goods.Apply(plr)

		plr.UpdateHeroToClient(req.Id)
	}()

	plr.SendPacket(protocol.MSG_SC_HeroLevelupResponse, res)

	if lv_old != lv_new {
		event.Fire(constant.Evt_Hero_LevelUp, plr, req.Id, lv_old, lv_new)
	}
}

func handler_HeroRefineRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.HeroRefineRequest{}
	res := &msg.HeroRefineResponse{}
	proto.Unmarshal(packet.Data, req)

	res.ErrorCode = ec.OK

	func() {

		hero := plr.GetHero(req.Id)
		if hero == nil {
			res.ErrorCode = ec.Hero_Not_Activated
			return
		}

		goods := app.NewItemProxy(constant.ItemLog_HeroRefine).SetArgs(hero.Id)

		if hero.RefineSuper {
			if req.Flag == 1 {
				conf := config.RefineSuperConf.Query(hero.RefineLv + 1)
				if conf == nil {
					res.ErrorCode = ec.Level_Exceed
					return
				}
				goods.Sub(constant.ItemID_RefineSuper, uint64(conf.Count))
				if !goods.Enough(plr) {
					res.ErrorCode = ec.Item_Not_Enough
					return
				}
				goods.Apply(plr)
				if rand.RandomHitn(int(conf.Prob), 100) {
					hero.RefineLv++
					res.Result = 1
				} else {
					hero.RefineLv = 0
				}
			} else {
				hero.RefineLv = 0
				hero.RefineTimes = 0
				hero.RefineSuper = false
				conf := config.RefineNormalConf.Query(hero.RefineLv + 1)
				if conf == nil {
					res.ErrorCode = ec.Level_Exceed
					return
				}
				goods.Sub(constant.ItemID_RefineNormal, uint64(conf.Count))
				if !goods.Enough(plr) {
					res.ErrorCode = ec.Item_Not_Enough
					return
				}
				goods.Apply(plr)
				if rand.RandomHitn(int(conf.Prob+hero.RefineTimes*2), 100) {
					hero.RefineLv++
					hero.RefineTimes = 0
					res.Result = 1
				} else {
					hero.RefineTimes++
				}
			}
		} else {
			if req.Flag == 1 {
				hero.RefineLv = 0
				hero.RefineTimes = 0
				hero.RefineSuper = true
				conf := config.RefineSuperConf.Query(hero.RefineLv + 1)
				if conf == nil {
					res.ErrorCode = ec.Level_Exceed
					return
				}
				goods.Sub(constant.ItemID_RefineSuper, uint64(conf.Count))
				if !goods.Enough(plr) {
					res.ErrorCode = ec.Item_Not_Enough
					return
				}
				goods.Apply(plr)
				if rand.RandomHitn(int(conf.Prob), 100) {
					hero.RefineLv++
					res.Result = 1
				} else {
					hero.RefineLv = 0
				}
			} else {
				conf := config.RefineNormalConf.Query(hero.RefineLv + 1)
				if conf == nil {
					res.ErrorCode = ec.Level_Exceed
					return
				}
				goods.Sub(constant.ItemID_RefineNormal, uint64(conf.Count))
				if !goods.Enough(plr) {
					res.ErrorCode = ec.Item_Not_Enough
					return
				}
				goods.Apply(plr)
				if rand.RandomHitn(int(conf.Prob+hero.RefineTimes*2), 100) {
					hero.RefineLv++
					hero.RefineTimes = 0
					res.Result = 1
				} else {
					hero.RefineTimes++
				}
			}
		}
	}()

	if res.ErrorCode == ec.OK {
		plr.UpdateHeroToClient(req.Id)
	}

	plr.SendPacket(protocol.MSG_SC_HeroRefineResponse, res)
}
