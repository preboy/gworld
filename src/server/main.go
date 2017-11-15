package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

import (
	"core/log"
	"core/schedule"
	"core/timer"
	"server/cmd"
	"server/net_mgr"
)

func main() {

	fmt.Println("server start ...")

	log.Start("GameServer")
	timer.Start()
	schedule.Start()
	net_mgr.Start()

	fmt.Println("server running ...")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \r\n\t")
		if strings.Compare(text, "quit") == 0 {
			break
		} else {
			cmd.ParseCommand(&text)
		}
	}

	fmt.Println("server closing")

	net_mgr.Stop()
	schedule.Start()
	timer.Stop()
	log.Stop()

	fmt.Println("server closed")
}
