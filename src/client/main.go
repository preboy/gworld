package main

import (
	"client/cmd"
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

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \r\n\t")
		if strings.Compare(text, "quit") == 0 {
			break
		} else {
			cmd.ParseCommand(&text)
		}
	}

	<-quit

	net_mgr.Stop()
}
