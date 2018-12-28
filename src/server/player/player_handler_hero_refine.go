package player

import (
	"core/math"
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
	"server/app"
	"server/config"
	"server/constant"
)

func init() {
	register_handler(protocol.MSG_CS_HeroRefine, handler_hero_refine)
}

func handler_hero_refine(plr *Player, packet *tcp.Packet) {
	req := msg.HeroRefineRequest{}
	res := msg.HeroRefineResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.ErrorCode = ec.OK

	func() {

		hero := plr.GetHero(req.HeroId)
		if hero == nil {
			res.ErrorCode = ec.Hero_Not_Activated
			return
		}

		goods := app.NewItemProxy(protocol.MSG_CS_HeroRefine)

		if hero.RefineSuper {
			if req.Super == 1 {
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
				if math.RandomHitn(int(conf.Prob), 100) {
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
				if math.RandomHitn(int(conf.Prob+hero.RefineTimes*2), 100) {
					hero.RefineLv++
					hero.RefineTimes = 0
					res.Result = 1
				} else {
					hero.RefineTimes++
				}
			}
		} else {
			if req.Super == 1 {
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
				if math.RandomHitn(int(conf.Prob), 100) {
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
				if math.RandomHitn(int(conf.Prob+hero.RefineTimes*2), 100) {
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
		plr.UpdateHeroToClient(req.HeroId)
	}
	plr.SendPacket(protocol.MSG_SC_HeroRefine, &res)
}
