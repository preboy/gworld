package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"gworld/core/log"
	"gworld/core/rand"
	"gworld/core/schedule"
	"gworld/core/timer"
	"gworld/core/utils"
	"gworld/core/wordsfilter"
	"gworld/core/work"
	"gworld/game/app"
	"gworld/game/cmd"
	"gworld/game/config"
	"gworld/game/dbmgr"
	"gworld/game/loop"
	"gworld/game/modules/act"
	"gworld/game/modules/websvr"
	"gworld/game/netmgr"
	"gworld/game/player"
	"gworld/game/world"

	_ "gworld/game/modules/preloader"
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

	// init server
	rand.SetSeed()

	if !app.LoadConfig("config.json", *arg_svr) {
		log.Error("app.LoadServerConfig: Failed")
		log.Stop()
		return
	}

	dbmgr.Open(app.GetGameConfig().DBGame, app.GetGameConfig().DBStat, app.GetConfig().Common.DBCenter)

	if !app.LoadServerData() {
		log.Stop()
		return
	}

	timer.Start()
	schedule.Start()
	work.Start(4)

	if wordsfilter.Load("./config/filter.txt") != nil {
		log.Error("wordsfilter.Load failed !")
		return
	}

	config.LoadAll(true)

	world.Start()

	player.LoadData()

	act.Open()

	loop.Start()

	netmgr.Start()

	websvr.Start()

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

	log.Info("netmgr stopping ...")
	netmgr.Stop()

	act.Close()

	log.Info("loop stopping ...")
	loop.Stop()

	log.Info("work stopping ...")
	work.Stop()

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
