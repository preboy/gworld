package main

import (
	"flag"
	"os"
	"syscall"

	"gworld/core/log"
	"gworld/core/utils"
	"gworld/ddz/lobby"
	"gworld/ddz/loop"
	"gworld/ddz/netmgr"
	"gworld/ddz/player"
)

var (
	quit = make(chan bool)
)

// ----------------------------------------------------------------------------
// main

func main() {
	arg_log := flag.String("log", "ddz.log", "log file name")
	flag.Parse()

	log.Start(*arg_log)
	log.Info("DDZ starting ...")

	utils.RegisterSignalHandler(func(sig os.Signal) {
		log.Warning("signal catched: %v", sig)
		if sig == syscall.SIGHUP {
			// reserved
		} else {
			close(quit)
		}
	})

	ddz_init()

	loop.Run()
	<-quit

	ddz_release()

	log.Info("DDZ closed !!!")
}

// ----------------------------------------------------------------------------
// local

func ddz_init() {
	lobby.Init()
	player.Init()
	netmgr.Init()
}

func ddz_release() {
	netmgr.Release()
	player.Release()
	lobby.Release()
}
