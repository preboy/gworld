package cmd

import (
	"fmt"
	"strings"

	"core/utils"
	"server/app"
	"server/battle"
	"server/config"
)

func ParseCommand(cmd string) {
	var args []string
	for _, v := range strings.Split(cmd, " ") {
		arg := strings.Trim(v, " ,\t")
		if arg != "" {
			args = append(args, arg)
		}
	}

	len_args := len(args)
	if len_args == 0 {
		return
	}

	switch args[0] {
	case "p":
		panic("self")

	case "b":
		a := app.CreatureTeamToBattleTroop(1)
		d := app.CreatureTeamToBattleTroop(2)
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

	case "load":
		if len_args > 1 {
			config.LoadOne(args[1])
		} else {
			fmt.Println("load config_name")
		}

	default:
		fmt.Println("unknown command !!!")
	}
}
