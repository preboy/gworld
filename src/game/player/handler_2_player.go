package player

import (
	"strings"

	"core/log"
	"core/tcp"
	"core/wordsfilter"
	"game/app"
	"game/config"
	"game/constant"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
)

func handler_PlayerDataRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.PlayerDataRequest{}
	proto.Unmarshal(packet.Data, req)

	res := plr.data.ToMsg()
	plr.SendPacket(protocol.MSG_SC_PlayerDataResponse, res)
}

func handler_GMCommandRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.GMCommandRequest{}
	res := &msg.GMCommandResponse{}
	proto.Unmarshal(packet.Data, req)

	var args []string
	for _, s := range strings.Split(req.Command, " ") {
		args = append(args, strings.Trim(s, ", "))
	}

	if len(args) > 0 {
		res.Result = plr.on_gm_command(args)
	}
	plr.SendPacket(protocol.MSG_SC_GMCommandResponse, res)
}

func handler_UseItemRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.UseItemRequest{}
	res := &msg.UseItemResponse{}
	proto.Unmarshal(packet.Data, req)

	res.Result = ec.OK

	func() {

		ItemID := req.Id
		ItemCt := req.Cnt

		// 是否存在此种道具
		ip := config.ItemProtoConf.Query(ItemID)
		if ip == nil {
			res.Result = ec.Item_Not_Enough
			return
		}

		// 道具数量是否足够
		goods := app.NewItemProxy(constant.ItemLog_UseItem).SetArgs(ItemID, ItemCt)
		goods.Sub(ItemID, uint64(ItemCt))
		if !goods.Enough(plr) {
			res.Result = ec.Item_Not_Enough
			return
		}

		// 是否可使用
		if ip.UseType == 0 {
			res.Result = ec.Item_Unusable
			return
		}

		switch ip.UseType {
		case 1: // 兑换道具
			{
				if ip.Param1 != 0 && ip.Param2 != 0 {
					goods.Add(uint32(ip.Param1), uint64(uint32(ip.Param2)*ItemCt))
				}
			}
		default:
			log.Warning("Invalid ITEM UseType: %v-%v", ItemID, ip.UseType)
		}

		goods.Apply(plr)
	}()

	plr.SendPacket(protocol.MSG_SC_UseItemResponse, res)
}

func handler_MarketBuyRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.MarketBuyRequest{}
	res := &msg.MarketBuyResponse{}
	proto.Unmarshal(packet.Data, req)

	res.ErrorCode = ec.Failed

	func() {
		// 检测包裹道具是否足够
		conf := config.MarketConf.Query(req.Index)
		if conf == nil {
			return
		}

		proxy := app.NewItemProxy(constant.ItemLog_MarketBuy)

		for _, item := range conf.Src {
			proxy.Sub(item.Id, uint64(item.Cnt)*req.Count)
		}

		if !proxy.Enough(plr) {
			res.ErrorCode = ec.Item_Not_Enough
			return
		}

		for _, item := range conf.Dst {
			proxy.Add(item.Id, uint64(item.Cnt)*req.Count)
		}

		proxy.Apply(plr)

		res.ErrorCode = ec.OK
	}()

	plr.SendPacket(protocol.MSG_SC_MarketBuyResponse, res)
}

func handler_ChangeNameRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.ChangeNameRequest{}
	res := &msg.ChangeNameResponse{}
	proto.Unmarshal(packet.Data, req)

	func() {
		Name := strings.TrimSpace(req.Name)

		if wordsfilter.IsPunctuation(Name) {
			res.ErrorCode = ec.NamePunctuation
			return
		}

		if wordsfilter.IsSensitive(Name) {
			res.ErrorCode = ec.NameSensitive
			return
		}

		// 长度
		if len(Name) == 0 || len(Name) > constant.MaxNameLength {
			res.ErrorCode = ec.NameLengthErr
			return
		}

		if Name == plr.GetName() {
			res.ErrorCode = ec.NameSame
			return
		}

		// TODO 未作冲突检测

		plr.SetName(Name)

		res.ErrorCode = ec.OK
	}()

	plr.SendPacket(protocol.MSG_SC_ChangeNameResponse, res)
}
