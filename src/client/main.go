package main

import (
	"client/net_mgr"
)

var (
	quit chan bool
)

func init() {
	quit = make(chan bool)
}

func main() {

	net_mgr.Start()

	<-quit

	net_mgr.Stop()
}
