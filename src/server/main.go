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
	"core/work_service"
	"server/cmd"
	"server/db_mgr"
	"server/game"
	"server/game/config"
	"server/net_mgr"
	"server/player"
	"server/server"
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
	work_service.Start(4)

	config.Load()

	player.LoadData()

	main_thread := server.NewServer()
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

	log.Info("net_mgr stopping ...")
	net_mgr.Stop()

	log.Info("main_thread stopping ...")
	main_thread.Stop()

	log.Info("work_service stopping ...")
	work_service.Stop()

	log.Info("player stopping ...")
	player.SaveData()

	log.Info("save server data ...")
	game.SaveServerData()

	log.Info("schedule stopping ...")
	schedule.Stop()

	log.Info("timer stopping ...")
	timer.Stop()

	log.Info("db_mgr stopping ...")
	db_mgr.Close()

	fmt.Println("server closed")
	log.Stop()
}
