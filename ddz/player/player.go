package player

import (
	"gworld/core/tcp"
	"gworld/ddz/loop"
)

var (
	_plrs = map[string]*Player{}
)

// ----------------------------------------------------------------------------
// init

func init() {
	loop.Register(func() {
		for _, plr := range _plrs {
			plr.OnUpdate()
		}
	})
}

// ----------------------------------------------------------------------------
// export

func Init() {
}

func Release() {
}

func NewPlayer(pid string) *Player {
	plr := &Player{
		PID: pid,
	}

	return plr
}

// ----------------------------------------------------------------------------
// Player

type Player struct {
	PID string
}

func (self *Player) OnLogin() {
	_plrs[self.PID] = self
}

func (self *Player) OnLogout() {
	delete(_plrs, self.PID)
}

func (self *Player) OnUpdate() {
}

func (self *Player) OnPacket(packet *tcp.Packet) {

}
