package main

import (
	"fmt"

	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/loop"
	"gworld/ddz_ai/args"
	"gworld/ddz_ai/netmgr"
)

func main() {
	if !args.Parse() {
		fmt.Println("Args parse failed")
		return
	}

	var logname = "ddz_ai_"
	logname += args.NickName
	logname += ".log"

	log.Start(logname)
	defer log.Stop()

	log.Info("boot with %v %v", args.NickName, args.MatchName)

	netmgr.Init()

	loop.Run()
	block.Wait()

	netmgr.Init()
}
