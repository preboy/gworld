package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_Notice, handler_player_notice)
}

func handler_player_notice(plr *Player, packet *tcp.Packet) {
	req := msg.NoticeRequest{}
	res := msg.NoticeResponse{}
	proto.Unmarshal(packet.Data, &req)

	res.Flag = req.Flag
	res.Notice = req.Notice

	plr.SendPacket(protocol.MSG_SC_Notice, &res)
}
