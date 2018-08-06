package cmd

import (
	"core/utils"
	"fmt"
	"server/game"
	"server/game/battle"
)

func ParseCommand(cmd *string) {
	switch *cmd {
	case "p":
		panic("self")

	case "b":
		a := game.CreatureTeamToBattleTroop(1)
		d := game.CreatureTeamToBattleTroop(2)
		b := battle.NewBattle(a, d)
		b.Calc()
		r := b.GetResult()
		fmt.Println("battle test:", a)
		fmt.Println("battle test:", d)
		fmt.Println("battle result:", r)

	case "start_prof":
		fmt.Println(utils.StartPprof("prof"))
	case "close_prof":
		fmt.Println(utils.ClosePprof())

	default:
		fmt.Println("unknown command !!!")
	}
}
