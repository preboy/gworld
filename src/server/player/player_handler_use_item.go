package player

import (
	"core/log"
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/err_code"
	"public/protocol"
	"public/protocol/msg"
	"server/game/config"
)

func init() {
	register_handler(protocol.MSG_CS_UseItem, handler_use_item)
}

func handler_use_item(plr *Player, packet *tcp.Packet) {
	req := msg.UseItemRequest{}
	res := msg.UseItemResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.Result = err_code.ERR_OK

	func() {

		ItemID := req.Id
		ItemCt := req.Cnt

		// 是否存在此种道具
		ip := config.GetItemProto(ItemID)
		if ip == nil {
			res.Result = err_code.ERR_UNKNOWN_ITEM
			return
		}

		// 道具数量是否足够
		goods := NewItemProxy(protocol.MSG_CS_UseItem)
		goods.Sub(ItemID, uint64(ItemCt))
		if !goods.Enough(plr) {
			res.Result = err_code.ERR_ITEM_NOT_ENOUGH
			return
		}

		// 是否可使用
		if ip.UseType == 0 {
			res.Result = err_code.ERR_ITEM_UNUSABLE
			return
		}

		switch ip.UseType {
		case 1: // 兑换道具
			{
				if ip.Param1 != 0 && ip.Param2 != 0 {
					goods.Add(uint32(ip.Param1), uint64(uint32(ip.Param2)*ItemCt))
				}
			}
		case 2: // 给英雄使用经验丹
			{
				HeroID := uint32(req.Arg1)
				hero := plr.GetHero(HeroID)
				if hero == nil {
					res.Result = err_code.ERR_ITEM_INVALID_HERO
					return
				}
				hero.AddExp(uint32(ip.Param1) * ItemCt)
			}
		default:
			log.Warning("Invalid ITEM UseType: %v-%v", ItemID, ip.UseType)
		}

		goods.Apply(plr)
	}()

	plr.SendPacket(protocol.MSG_SC_UseItem, &res)
}
