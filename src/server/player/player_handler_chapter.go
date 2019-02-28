package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_ChapterInfo, handler_player_ChapterInfo)
	register_handler(protocol.MSG_CS_ChapterFighting, handler_player_ChapterFighting)
	register_handler(protocol.MSG_CS_ChapterRewards, handler_player_ChapterRewards)
	register_handler(protocol.MSG_CS_ChapterLoot, handler_player_ChapterLoot)
}

func handler_player_ChapterInfo(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterInfoRequest{}
	res := &msg.ChapterInfoResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.Getchapter().ChapterInfo(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterInfo, res)
}

func handler_player_ChapterFighting(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterFightingRequest{}
	res := &msg.ChapterFightingResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.Getchapter().ChapterFighting(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterFighting, res)
}

func handler_player_ChapterRewards(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterRewardsRequest{}
	res := &msg.ChapterRewardsResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.Getchapter().ChapterRewards(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterRewards, res)
}

func handler_player_ChapterLoot(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterLootRequest{}
	res := &msg.ChapterLootResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.Getchapter().ChapterLoot(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterLoot, res)
}
