package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

import (
	"clients"
	"utils"
)

func main() {

	fmt.Println("server start ...")

	clients.StartListen()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \r\n\t")

		if strings.Compare(text, "quit") == 0 {
			break
		} else {
			utils.ParseCommand(&text)
		}
	}

	fmt.Println("server closing")

	clients.EndListen()

	fmt.Println("server closed")
}
