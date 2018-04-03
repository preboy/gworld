package player

import (
	"core/tcp"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
)

func init() {
	register_handler(protocol.MSG_CS_PlayerData, handler_player_data)
}

func handler_player_data(plr *Player, packet *tcp.Packet) {
	req := msg.PlayerDataRequest{}
	res := msg.PlayerDataResponse{}
	proto.Unmarshal(packet.Data, &req)
	// TODO something

	data := plr.GetData()

	res.Acct = data.Acct
	res.Name = data.Name
	res.Pid = data.Pid
	res.Sid = plr.sid
	res.Id = req.Id
	res.Level = data.Level
	res.VipLevel = data.VipLevel
	res.Male = data.Male

	for id, cnt := range data.Items {
		res.Items = append(res.Items, &msg.Item{
			Flag: 0,
			Id:   id,
			Cnt:  int64(cnt),
		})
	}

	for _, hero := range data.Heros {
		res.Heros = append(res.Heros, hero.ToMsg())
	}

	plr.SendPacket(protocol.MSG_SC_PlayerData, &res)
}
