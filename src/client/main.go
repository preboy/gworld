package main

import (
	"client/cmd"
	"client/net_mgr"
)

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

	net_mgr.Stop()
}
