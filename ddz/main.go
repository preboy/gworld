package main

import (
	"flag"
	"os"
	"syscall"

	"gworld/core/block"
	"gworld/core/log"
	"gworld/core/rand"
	"gworld/core/utils"
	"gworld/ddz/config"
	"gworld/ddz/gambler"
	"gworld/ddz/lobby"
	"gworld/ddz/loop"
	"gworld/ddz/netmgr"
	"gworld/ddz/referee"
)

// ----------------------------------------------------------------------------
// main

func main() {
	arg_log := flag.String("log", "ddz.log", "log file name")
	flag.Parse()

	log.Start(*arg_log)
	defer log.Stop()

	log.Info("DDZ starting ...")

	utils.RegisterSignalHandler(func(sig os.Signal) {
		log.Warning("signal catched: %v", sig)
		if sig == syscall.SIGHUP {
			// reserved
		} else {
			block.Done()
		}
	})

	if err := config.Load(); err != nil {
		panic(err)
	}

	// init server
	rand.SetSeed()

	ddz_init()

	loop.Run()
	block.Wait()

	ddz_release()

	log.Info("DDZ closed !!!")
}

// ----------------------------------------------------------------------------
// local

func ddz_init() {
	// init modules
	lobby.Init()
	gambler.Init()
	referee.Init()
	netmgr.Init()
}

func ddz_release() {
	// release modules
	netmgr.Release()
	gambler.Release()
	referee.Release()
	lobby.Release()
}
