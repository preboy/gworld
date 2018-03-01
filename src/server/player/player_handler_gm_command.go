package player

import (
	"core/tcp"
	"core/utils"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
	"strings"
)

func init() {
	register_handler(protocol.MSG_CS_GMCommand, handler_gm_command)
}

func handler_gm_command(plr *Player, packet *tcp.Packet) {
	req := msg.GMCommandRequest{}
	res := msg.GMCommandResponse{}
	proto.Unmarshal(packet.Data, &req)

	var args []string
	for _, s := range strings.Split(req.Command, " ") {
		args = append(args, strings.Trim(s, ", "))
	}

	if len(args) > 0 {
		res.Result = plr.on_gm_command(args)
	}

	plr.SendPacket(protocol.MSG_SC_GMCommand, &res)
}

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
			ip := NewItemProxy()
			item := strings.Split(args[i], "|")
			id := utils.Atou32(item[0])
			ct := utils.Atou32(item[1])
			ip.Add(id, uint64(ct))
			ip.Apply(self)
			println("curr:", item[0], self.GetItemCnt(id))
		}
	default:
		println("unknown command:", args[0])
		return 0
	}
	return 1
}
