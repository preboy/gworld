package player

import ()

type ISession interface {
	Send(data []byte)
	SetPlayer(*Player)
}

func (self *Player) SetSession(s ISession) {
	self.s = s
}

func (self *Player) Send(data []byte) {
	self.s.Send(data)
}
