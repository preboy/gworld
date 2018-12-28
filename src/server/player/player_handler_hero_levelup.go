package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
	"server/app"
	"server/config"
)

func init() {
	register_handler(protocol.MSG_CS_HeroLevelup, handler_hero_levelup)
}

func handler_hero_levelup(plr *Player, packet *tcp.Packet) {
	req := msg.HeroLevelupRequest{}
	res := msg.HeroLevelupResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.ErrorCode = ec.OK

	var lv_old, lv_new uint32

	func() {
		hero := plr.GetHero(req.HeroId)
		if hero == nil {
			res.ErrorCode = ec.Hero_Not_Activated
			return
		}

		// 可否升到下一级
		if config.HeroConf.Query(hero.Id, hero.Level+1) == nil {
			return
		}

		conf := config.HeroConf.Query(hero.Id, hero.Level)

		// 道具数量是否足够
		goods := app.NewItemProxy(protocol.MSG_CS_HeroLevelup)
		for _, v := range conf.Needs {
			goods.Sub(v.Id, v.Cnt)
		}
		if !goods.Enough(plr) {
			res.ErrorCode = ec.Item_Not_Enough
			return
		}

		lv_old = hero.Level
		hero.Level++
		lv_new = hero.Level
		goods.Apply(plr)
		plr.UpdateHeroToClient(req.HeroId)

	}()

	plr.SendPacket(protocol.MSG_SC_HeroLevelup, &res)

	if lv_old != lv_new {
		plr.OnPlayerLevelup(lv_old, lv_new)
	}
}
