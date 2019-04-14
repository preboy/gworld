package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"core/log"
	"core/schedule"
	"core/timer"
	"core/utils"
	"core/work_service"
	"game/app"
	"game/cmd"
	"game/config"
	"game/dbmgr"
	"game/loop"
	"game/microsvr"
	"game/modules/act"
	"game/netmgr"
	"game/player"
	"game/world"

	_ "game/microsvr/preloader"
	_ "game/modules/preloader"
)

var (
	quit = make(chan bool)
)

func main() {
	arg_svr := flag.String("svr", "game1", "server id")
	arg_log := flag.String("log", "game1.log", "log file name")
	flag.Parse()

	log.Start(*arg_log)
	log.Info("server start ...")

	utils.RegisterSignalHandler(func(sig os.Signal) {
		if sig == syscall.SIGHUP {
			fmt.Println("signal catched: syscall.SIGHUP")
		} else {
			close(quit)
			os.Stdin.Close()
		}
	})

	if !app.LoadConfig("config.json", *arg_svr) {
		log.Error("app.LoadServerConfig: Failed")
		log.Stop()
		return
	}

	dbmgr.Open(app.GetGameConfig().DBGame, app.GetGameConfig().DBStat)

	app.LoadServerData()

	timer.Start()
	schedule.Start()
	work_service.Start(4)

	config.LoadAll(true)

	world.Start()

	player.LoadData()

	microsvr.Start()

	act.Open()

	loop.Start()

	netmgr.Start()

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
			cmd.ParseCommand(text)
		}
	}

	<-quit

	log.Info("server stopping ...")

	log.Info("net_mgr stopping ...")
	net_mgr.Stop()

	act.Close()

	log.Info("micro svrs stopping ...")
	microsvr.Stop()

	log.Info("loop stopping ...")
	loop.Stop()

	log.Info("work_service stopping ...")
	work_service.Stop()

	log.Info("player stopping ...")
	player.SaveData()

	log.Info("world stopping ...")
	world.Stop()

	log.Info("save server data ...")
	app.SaveServerData()

	log.Info("schedule stopping ...")
	schedule.Stop()

	log.Info("timer stopping ...")
	timer.Stop()

	log.Info("dbmgr stopping ...")
	dbmgr.Close()

	fmt.Println("server closed")
	log.Stop()
}
