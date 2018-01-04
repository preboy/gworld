package main

import (
	_ "bufio"
	"fmt"
	"os"
	_ "strings"
	"syscall"
)

import (
	"core/db"
	"core/log"
	"core/schedule"
	"core/timer"
	"core/utils"
	_ "server/cmd"
	"server/game"
	"server/game/config"
	"server/net_mgr"
)

var (
	quit = make(chan bool)
)

func main() {
	log.Start("GameServer")

	log.Info("server start ...")

	if !game.Init() {
		log.Error("Fail on game.Init")
		log.Stop()
		// time.Sleep(100 * time.Microsecond)
		return
	}

	utils.RegisterSignalHandler(func(sig os.Signal) {
		if sig == syscall.SIGHUP {
			fmt.Println("signal catched: syscall.SIGHUP")
		} else {
			close(quit)
		}
	})

	db.Open(game.GetServerConfig().DBAddr)

	timer.Start()
	schedule.Start()
	net_mgr.Start()

	config.Load()

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	text, _ := reader.ReadString('\n')
	// 	text = strings.Trim(text, " \r\n\t")
	// 	if strings.Compare(text, "quit") == 0 {
	// 		break
	// 	} else {
	// 		cmd.ParseCommand(&text)
	// 	}
	// }

	log.Info("server running ...")
	<-quit
	log.Info("server stopping ...")

	net_mgr.Stop()
	schedule.Start()
	timer.Stop()

	db.Close()

	fmt.Println("server closed")

	log.Stop()
}
