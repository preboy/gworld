package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/err_code"
	"public/protocol"
	"public/protocol/msg"
	"server/game/config"
)

func init() {
	register_handler(protocol.MSG_CS_HeroLevelup, handler_hero_levelup)
}

func handler_hero_levelup(plr *Player, packet *tcp.Packet) {
	req := msg.HeroLevelupRequest{}
	res := msg.HeroLevelupResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.ErrorCode = err_code.ERR_OK

	func() {
		hero := plr.GetHero(req.HeroId)
		if hero == nil {
			res.ErrorCode = err_code.ERR_INVALID_HERO
			return
		}

		// 可否升到下一级
		if config.GetHeroProto(hero.Id, hero.Level+1) == nil {
			return
		}

		conf := config.GetHeroProto(hero.Id, hero.Level)

		// 道具数量是否足够
		goods := NewItemProxy(protocol.MSG_CS_HeroLevelup)
		for _, v := range conf.Needs {
			goods.Sub(v.Id, uint64(v.Cnt))
		}
		if !goods.Enough(plr) {
			res.ErrorCode = err_code.ERR_ITEM_NOT_ENOUGH
			return
		}

		hero.Level++
		goods.Apply(plr)
		plr.UpdateHeroToClient(req.HeroId)
	}()

	plr.SendPacket(protocol.MSG_SC_HeroLevelup, &res)
}
