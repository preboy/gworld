package main

import (
	"gworld/core/log"
	"gworld/ddz/loop"
	"gworld/ddz_rr/netmgr"
)

var (
	quit = make(chan bool)
)

func main() {
	log.Start("ddz_rr.log")

	netmgr.Init()

	loop.Run()

	<-quit

	netmgr.Release()
}
