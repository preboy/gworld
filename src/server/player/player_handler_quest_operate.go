package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/ec"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_QuestOp, handler_player_operate)
}

func handler_player_operate(plr *Player, packet *tcp.Packet) {
	req := msg.QuestOpRequest{}
	proto.Unmarshal(packet.Data, &req)

	res := msg.QuestOpResponse{
		Id: req.Id,
		Op: req.Op, // 1:接受 2:放弃 3:提交  4:完成
		R:  req.R,
	}

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

	plr.SendPacket(protocol.MSG_SC_QuestOp, &res)
}
