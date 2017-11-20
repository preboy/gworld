package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_PlayerData, handler_player_data)
}

func handler_player_data(plr *Player, packet *tcp.Packet) {
	req := msg.PlayerDataRequest{}
	res := msg.PlayerDataResponse{}
	proto.Unmarshal(packet.Data, &req)
	// TODO something

	pd := plr.GetData()

	res.Acct = pd.Acct
	res.Name = pd.Name
	res.Pid = pd.Pid
	res.Sid = plr.sid
	res.Id = req.Id

	plr.SendPacket(packet.Opcode, &res)
}
