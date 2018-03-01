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
	proto.Unmarshal(packet.Data, &req)
	plr.SendNotice(req.Notice, req.Flag)
}
