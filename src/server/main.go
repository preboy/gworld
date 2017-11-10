package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

import (
	"core/log"
	"server/cmd"
	"server/net_mgr"
)

func main() {

	log.Start("GameServer")

	fmt.Println("server start ...")

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

	log.Stop()

	fmt.Println("server closed")
}
