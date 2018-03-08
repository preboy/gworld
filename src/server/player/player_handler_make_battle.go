package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_MakeBattle, handler_player_make_battle)
}

func handler_player_make_battle(plr *Player, packet *tcp.Packet) {
	req := msg.MakeBattleRequest{}
	res := msg.MakeBattleResponse{}

	proto.Unmarshal(packet.Data, &req)

	plr.SendPacket(protocol.MSG_SC_MakeBattle, &res)
}
