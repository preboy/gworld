package session

import (
	"fmt"
)

import (
	"github.com/gogo/protobuf/proto"
)

import (
	"core/tcp"
	"public/err_code"
	"public/protocol"
	"public/protocol/msg"
)

var (
	stages [10]StageFunction
	idx    int
)

type StageFunction struct {
	OnEnter  func(s *Session)
	OnLeave  func(s *Session)
	OnUpdate func(s *Session)
	OnPacket func(s *Session, packet *tcp.Packet)
	OnTimer  func(s *Session, id uint64)
}

type StageLogin struct {
	send   bool
	verify bool
}

func init() {

	idx = 0
	stages[idx].OnEnter = func(s *Session) {
	}
	stages[idx].OnLeave = func(s *Session) {
	}
	stages[idx].OnUpdate = func(s *Session) {
	}
	stages[idx].OnPacket = func(s *Session, packet *tcp.Packet) {
	}
	stages[idx].OnTimer = func(s *Session, id uint64) {
	}

	//------------------------------------------------------------------
	idx = 1 // login
	stages[idx].OnEnter = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnEnter")
		req := &msg.LoginRequest{}
		req.Acct = "test341"
		req.Pass = "1"
		s.SendPacket(protocol.MSG_LOGIN, req)

		s.stage_data = &StageLogin{
			send:   true,
			verify: false,
		}
	}

	stages[idx].OnLeave = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnLeave")

	}

	stages[idx].OnUpdate = func(s *Session) {
		// fmt.Println("Stage:", s.stage_id, "OnUpdate")
		sd := s.stage_data.(*StageLogin)
		if sd.verify {
			fmt.Println("Login OK")
			Next(s)
		}
	}

	stages[idx].OnPacket = func(s *Session, packet *tcp.Packet) {
		if packet.Opcode == protocol.MSG_LOGIN {
			res := &msg.LoginResponse{}
			err := proto.Unmarshal(packet.Data, res)
			if err == nil {
				if res.ErrorCode == err_code.ERR_OK {
					sd := s.stage_data.(*StageLogin)
					sd.verify = true
				} else {
				}
			} else {
				fmt.Println("proto.Unmarshal ERROR", err)
			}
		}
	}

	stages[idx].OnTimer = func(s *Session, id uint64) {
	}

	//------------------------------------------------------------------
	idx = 2 // game
	stages[idx].OnEnter = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnEnter")
	}

	stages[idx].OnLeave = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnLeave")
	}

	stages[idx].OnUpdate = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnUpdate")
	}

	stages[idx].OnPacket = func(s *Session, packet *tcp.Packet) {
	}

	stages[idx].OnTimer = func(s *Session, id uint64) {
	}
}

func Next(s *Session) {
	stages[s.stage_id].OnLeave(s)
	s.stage_id++
	stages[s.stage_id].OnEnter(s)
}
