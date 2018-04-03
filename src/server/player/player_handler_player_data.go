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

	for id, hero := range data.Heros {
		_hero := &msg.Hero{
			Id:           id,
			Level:        hero.Level,
			RefineLv:     hero.RefineLv,
			RefineTimes:  hero.RefineTimes,
			RefineSuper:  hero.RefineSuper,
			Power:        hero.Power,
			Status:       hero.Status,
			LifePoint:    hero.LifePoint,
			LifePointMax: hero.LifePointMax,
		}

		for i := 0; i < 2; i++ {
			_hero.Active = append(_hero.Active, &msg.Skill{
				Id:    hero.Active[i].Id,
				Level: hero.Active[i].Level,
			})
		}

		for i := 0; i < 4; i++ {
			_hero.Passive = append(_hero.Passive, &msg.Skill{
				Id:    hero.Passive[i].Id,
				Level: hero.Passive[i].Level,
			})
		}

		res.Heros = append(res.Heros, _hero)
	}

	plr.SendPacket(protocol.MSG_SC_PlayerData, &res)
}
