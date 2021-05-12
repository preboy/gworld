package main

import (
	"gworld/core/block"
	"gworld/core/log"
	"gworld/ddz/loop"
	"gworld/ddz_rr/netmgr"
)

func main() {
	log.Start("ddz_rr.log")
	defer log.Stop()

	log.Info("I'am ddz_referee")

	netmgr.Init()

	loop.Run()

	block.Signal()

	netmgr.Release()
}
