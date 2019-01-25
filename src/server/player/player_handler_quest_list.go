package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_QuestList, handler_player_quest_list)
}

func handler_player_quest_list(plr *Player, packet *tcp.Packet) {
	req := msg.QuestListRequest{}
	proto.Unmarshal(packet.Data, &req)

	Quest := plr.GetData().Quest

	res := &msg.QuestListResponse{
		Quests: Quest.ToMsgs(),
	}

	plr.SendPacket(protocol.MSG_SC_QuestList, res)
}
