package cmd

import (
	"fmt"
)

func ParseCommand(cmd *string) {
	switch *cmd {
	case "":

	case "a":
		fmt.Println("you input a")

	default:
		fmt.Println("unknown command !!!")
	}
}
