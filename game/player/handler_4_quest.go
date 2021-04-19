package player

import (
	"gworld/core/tcp"
	"gworld/public/ec"
	"gworld/public/protocol"
	"gworld/public/protocol/msg"

	"github.com/gogo/protobuf/proto"
)

func handler_QuestListRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.QuestListRequest{}
	res := &msg.QuestListResponse{}
	proto.Unmarshal(packet.Data, req)

	res.Quests = plr.GetData().Quest.ToMsgs()

	plr.SendPacket(protocol.MSG_SC_QuestListResponse, res)
}

func handler_QuestOpRequest(plr *Player, packet *tcp.Packet) {
	req := &msg.QuestOpRequest{}
	res := &msg.QuestOpResponse{}
	proto.Unmarshal(packet.Data, req)

	res.Id = req.Id
	res.Op = req.Op // 1:接受 2:放弃 3:提交  4:完成
	res.R = req.R

	Quest := plr.GetData().Quest

	switch req.Op {
	case 1:
		res.ErrorCode = Quest.Accept(req.Id)
	case 2:
		res.ErrorCode = Quest.Cancel(req.Id)
	case 3:
		res.ErrorCode = Quest.Commit(req.Id, req.R)
	case 4:
		res.ErrorCode = Quest.Finish(req.Id)
	default:
		res.ErrorCode = ec.Failed
	}

	res.Quest = Quest.ToMsg(req.Id)

	plr.SendPacket(protocol.MSG_SC_QuestOpResponse, res)
}
