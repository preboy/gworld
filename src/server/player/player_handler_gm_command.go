package player

import (
	"core/tcp"
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
	for _, s := range strings.Split(req.Command, ",") {
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
	case "exp":
		println("exp")
	case "lv":
		println("lv")
	default:
		println("unknown command:", args[0])
		return 0
	}
	return 1
}
