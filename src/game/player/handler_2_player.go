package player

import (
	"fmt"
	"strings"

	"core/log"
	"core/tcp"
	"core/utils"
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
	res := &msg.PlayerDataResponse{}
	proto.Unmarshal(packet.Data, req)

	data := plr.GetData()

	res.Acct = data.Acct
	res.Name = data.Name
	res.Pid = data.Pid
	res.Sid = plr.sid
	res.Id = req.Id
	res.Level = data.Level
	res.VipLevel = data.VipLevel
	res.Male = data.Male

	for id, cnt := range data.Items {
		res.Items = append(res.Items, &msg.Item{
			Flag: 0,
			Id:   id,
			Cnt:  int64(cnt),
		})
	}

	for _, hero := range data.Heros {
		res.Heros = append(res.Heros, hero.ToMsg())
	}

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

// ----------------------------------------------------------------------------

func (self *Player) on_gm_command(args []string) int32 {
	fmt.Println("on_gm_command:", args)
	switch args[0] {
	case "save":
		self.Save()
	case "vip":
		if len(args) > 1 {
			val := utils.Atou32(args[1])
			self.data.VipLevel = val
			self.SendNotice("VipLevel: "+utils.U32toa(val), 0)
		}
	case "lv":
		if len(args) > 1 {
			val := utils.Atou32(args[1])
			self.data.Level = val
			self.SendNotice("Level: "+utils.U32toa(val), 0)
		}
	case "item":
		for i := 1; i < len(args); i++ {
			ip := app.NewItemProxy(constant.ItemLog_GM)
			item := strings.Split(args[i], "|")
			id := utils.Atou32(item[0])
			ct := utils.Atou32(item[1])
			ip.Add(id, uint64(ct))
			ip.Apply(self)
			println("curr:", item[0], self.GetItem(id))
		}
	case "hero":
		for i := 1; i < len(args); i++ {
			item := strings.Split(args[i], "|")
			id := utils.Atou32(item[0])
			hero := self.GetHero(id)
			if hero == nil {
				self.AddHero(id)
			} else {
				self.SendNotice("Hero: "+item[0]+" already exist", 0)
			}
		}
	default:
		println("unknown command:", args[0])
		return 0
	}
	return 1
}
