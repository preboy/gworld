package main

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/ddz/config"
	"gworld/ddz/loop"
	"gworld/ddz_rr/netmgr"
)

type session struct {
}

func (self *session) OnRecvPacket(packet *tcp.Packet) {

}
func (self *session) OnOpened() {
}
func (self *session) OnClosed() {
}

func (self *session) SendPacket(opcode uint16, data []byte) {

}

var (
	quit = make(chan bool)
)

func main() {
	log.Start("ddz_rr.log")

	if err := config.Load(); err != nil {
		panic(err)
	}

	netmgr.Init()

	loop.Run()

	<-quit

	netmgr.Release()
}
