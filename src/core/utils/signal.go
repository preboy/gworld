package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignalHandler(f func(os.Signal)) {
	go func() {
		fmt.Println("register signal complete")

		c := make(chan os.Signal, 10)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		for s := range c {
			f(s)
			if s != syscall.SIGHUP {
				break
			}
		}
	}()
}
