package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func handler_ChapterInfoRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterInfoRequest{}
	res := &msg.ChapterInfoResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.GetChapter().ChapterInfo(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterInfoResponse, res)
}

func handler_ChapterFightingRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterFightingRequest{}
	res := &msg.ChapterFightingResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.GetChapter().ChapterFighting(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterFightingResponse, res)
}

func handler_ChapterRewardsRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterRewardsRequest{}
	res := &msg.ChapterRewardsResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.GetChapter().ChapterRewards(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterRewardsResponse, res)
}

func handler_ChapterLootRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.ChapterLootRequest{}
	res := &msg.ChapterLootResponse{}
	proto.Unmarshal(packet.Data, req)

	plr.GetChapter().ChapterLoot(req, res)

	plr.SendPacket(protocol.MSG_SC_ChapterLootResponse, res)
}
