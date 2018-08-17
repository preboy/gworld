package player

import (
	"core/tcp"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"public/protocol"
	"public/protocol/msg"
	"server/game"
	"server/game/battle"
)

func init() {
	register_handler(protocol.MSG_CS_MakeBattle, handler_player_make_battle)
}

func handler_player_make_battle(plr *Player, packet *tcp.Packet) {
	req := msg.MakeBattleRequest{}
	res := msg.MakeBattleResponse{}

	proto.Unmarshal(packet.Data, &req)

	a := game.CreatureTeamToBattleTroop(req.Id)
	d := game.CreatureTeamToBattleTroop(2)
	b := battle.NewBattle(a, d)
	b.Calc()
	res.Result = b.ToMsg()

	plr.SendPacket(protocol.MSG_SC_MakeBattle, &res)

	fmt.Println("battle attacker:", a)
	fmt.Println("battle defender:", d)
	fmt.Println("battle result:", b.GetResult())
}
