package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

import (
	"core/db"
	"core/log"
	"core/schedule"
	"core/timer"
	"core/utils"
	"server/cmd"
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

	log.Info("server running ...")

	reader := bufio.NewReader(os.Stdin)

	// set debug flag
	if true {
		for {
			text, _ := reader.ReadString('\n')
			text = strings.Trim(text, " \r\n\t")
			if text == "" {
				continue
			}
			if strings.Compare(text, "quit") == 0 {
				close(quit)
				break
			} else {
				cmd.ParseCommand(&text)
			}
		}
	} else {
		<-quit
	}

	log.Info("server stopping ...")

	net_mgr.Stop()
	schedule.Start()
	timer.Stop()

	db.Close()

	fmt.Println("server closed")

	log.Stop()
}
