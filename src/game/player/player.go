package player

import (
	"time"

	"core/event"
	"game/constant"
	"public/protocol"
)

type Player struct {
	s      ISession
	sid    uint32 // slot id
	data   *PlayerData
	online bool
}

func NewPlayer() *Player {
	return &Player{}
}

// ============================================================================

func (self *Player) on_load() {
}

// ============================================================================
// private function

func (self *Player) Init() {
	self.data.Init(self)
	self.pursue()
}

// ============================================================================
// public function

func (self *Player) GetSid() int {
	return int(self.sid)
}

func (self *Player) Login() {
	self.online = true
	self.data.LoginTs = time.Now()
	self.data.LoginTimes++

	pid := self.data.Pid
	_plrs_online[pid] = self

	if self.data.LoginTimes == 1 {
		event.Fire(constant.Evt_Plr_LoginFirst, pid)
	}

	event.Fire(constant.Evt_Plr_Login, pid)

	// 发送玩家基本数据
	res := self.data.ToMsg()
	self.SendPacket(protocol.MSG_SC_PlayerDataResponse, res)

	if self.data.LoginTimes == 1 {
		self.AsyncSave()
	}
}

func (self *Player) Logout() {
	pid := self.data.Pid
	event.Fire(constant.Evt_Plr_Logout, pid)
	_plrs_online[pid] = nil

	self.s.Disconnect()
	self.s = nil
	self.online = false
	self.data.LogoutTs = time.Now()
}

func (self *Player) IsOnLine() bool {
	return self.online
}
