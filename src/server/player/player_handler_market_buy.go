package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/err_code"
	"public/protocol"
	"public/protocol/msg"
	"server/config"
)

func init() {
	register_handler(protocol.MSG_CS_MarketBuy, handler_player_market_buy)
}

func handler_player_market_buy(plr *Player, packet *tcp.Packet) {
	req := msg.MarketBuyRequest{}
	res := msg.MarketBuyResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.ErrorCode = err_code.ERR_FAILED

	func() {
		// 检测包裹道具是否足够
		conf := config.MarketConf.Query(req.Index)
		if conf == nil {
			return
		}

		proxy := NewItemProxy(protocol.MSG_CS_MarketBuy)

		for _, item := range conf.Src {
			proxy.Sub(item.Id, uint64(item.Cnt)*req.Count)
		}

		if !proxy.Enough(plr) {
			res.ErrorCode = err_code.ERR_ITEM_NOT_ENOUGH
			return
		}

		for _, item := range conf.Dst {
			proxy.Add(item.Id, uint64(item.Cnt)*req.Count)
		}

		proxy.Apply(plr)

		res.ErrorCode = err_code.ERR_OK
	}()

	plr.SendPacket(protocol.MSG_SC_MarketBuy, &res)
}
