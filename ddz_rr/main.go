package main

import (
	"gworld/core/loop"
	"gworld/core/tcp"
	"gworld/ddz/config"
	"gworld/game/loop"
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
	if err := config.Load(); err != nil {
		panic(err)
	}

	loop.Start()

	sess := &session{}

	{
		conn, err := tcp.Connect(config.Get().Addr4Referee)
		if err != nil {
			panic(err)
		}

		sock := tcp.NewSocket(conn, sess)
		sock.Start()
	}

	<-quit

}
