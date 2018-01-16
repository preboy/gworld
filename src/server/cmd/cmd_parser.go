package cmd

import (
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
		r := b.Calc()
		fmt.Println("battle test:", a)
		fmt.Println("battle test:", d)
		fmt.Println("battle result:", r)

	default:
		fmt.Println("unknown command !!!")
	}
}
