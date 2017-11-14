package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_LOGIN, handler_login)
}

func handler_login(plr *Player, packet *tcp.Packet) {
	req := msg.LoginRequest{}
	res := msg.LoginResponse{}
	proto.Unmarshal(packet.Data, &req)
	// TODO something

	plr.SendPacket(packet.Opcode, &res)
}
