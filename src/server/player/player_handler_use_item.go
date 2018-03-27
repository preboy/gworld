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
	register_handler(protocol.MSG_CS_UseItem, handler_use_item)
}

func handler_use_item(plr *Player, packet *tcp.Packet) {
	req := msg.UseItemRequest{}
	res := msg.UseItemResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.Result = err_code.ERR_OK

	func() {
		// 是否存在此种道具
		ip := config.GetItemProtoConf().GetItemProto(req.Id)
		if ip == nil {
			res.Result = err_code.ERR_UNKNOWN_ITEM
			return
		}
		// 是否可使用
		if ip.Usable != 1 {
			res.Result = err_code.ERR_ITEM_UNUSABLE
			return
		}
		// 无可用脚本ID
		script, ok := _item_scripts[ip.ScriptID]
		if !ok {
			res.Result = err_code.ERR_ITEM_INVALID_SCRIPTID
			return
		}
		// 道具数量是否足够
		goods := NewItemProxy(protocol.MSG_CS_UseItem)
		goods.Sub(req.Id, uint64(req.Cnt))
		if !goods.Enough(plr) {
			res.Result = err_code.ERR_ITEM_NOT_ENOUGH
			return
		}
		goods.Apply(plr)

		// 执行脚本
		script(plr, ip, req.Cnt)
	}()

	plr.SendPacket(protocol.MSG_SC_UseItem, &res)
}
