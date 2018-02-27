package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

import (
	"core/log"
	"core/schedule"
	"core/timer"
	"core/utils"
	"server/cmd"
	"server/db_mgr"
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

	utils.RegisterSignalHandler(func(sig os.Signal) {
		if sig == syscall.SIGHUP {
			fmt.Println("signal catched: syscall.SIGHUP")
		} else {
			close(quit)
			os.Stdin.Close()
		}
	})

	if !game.LoadServerConfig("config.json") {
		log.Error("game.LoadServerConfig: Failed")
		log.Stop()
		return
	}

	db_mgr.Open(game.GetServerConfig().DBAddr)

	game.LoadServerData()

	timer.Start()
	schedule.Start()
	net_mgr.Start()

	config.Load()

	main_thread := game.NewServer()
	main_thread.Start()

	log.Info("server running ...")

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Info("os.Stdin.Closed !")
			break
		}
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

	<-quit

	log.Info("server stopping ...")

	main_thread.Stop()

	net_mgr.Stop()
	schedule.Stop()
	timer.Stop()

	game.SaveServerData()

	db_mgr.Close()

	fmt.Println("server closed")

	log.Stop()
}
