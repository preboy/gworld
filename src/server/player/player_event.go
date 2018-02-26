package player

import (
	"fmt"
)

import (
	"core/event"
)

func (self *Player) FireEvent(evt *event.Event) {
	self.OnEvent(evt)
}

func (self *Player) OnEvent(evt *event.Event) {
	fmt.Println("Player.OnEvent:", evt)
}
