package main

import (
	"gworld/core/log"
	"gworld/ddz/loop"
	"gworld/ddz_ai/args"
	"gworld/ddz_ai/netmgr"
)

var (
	quit = make(chan bool)
)

func main() {
	log.Start("ddz_ai.log")
	defer log.Stop()

	log.Info("I'am ddz_ai")

	if !args.Parse() {
		log.Info("Args parse failed")
		return
	}

	log.Info("boot with %v %v", args.MatchID, args.NickName)

	netmgr.Init()

	loop.Run()

	<-quit

	netmgr.Init()
}
