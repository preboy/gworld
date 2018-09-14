package session

import (
	"github.com/gogo/protobuf/proto"

	"fmt"

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

type StageGame struct {
	ingame bool
	tid    uint64
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
		s.SendPacket(protocol.MSG_CS_LOGIN, req)

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
		if packet.Opcode == protocol.MSG_SC_LOGIN {
			res := &msg.LoginResponse{}
			err := proto.Unmarshal(packet.Data, res)
			if err == nil {
				if res.ErrorCode == err_code.ERR_OK {
					sd := s.stage_data.(*StageLogin)
					sd.verify = true
				} else {
					fmt.Println("LoginResponse, err_code = ", res.ErrorCode)
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
		req := &msg.EnterGameRequest{}
		s.SendPacket(protocol.MSG_CS_ENTER_GAME, req)
		s.stage_data = &StageGame{
			ingame: false,
		}
		fmt.Println("MSG_ENTER_GAME Request")
	}

	stages[idx].OnLeave = func(s *Session) {
		fmt.Println("Stage:", s.stage_id, "OnLeave")
	}

	stages[idx].OnUpdate = func(s *Session) {
		// fmt.Println("Stage:", s.stage_id, "OnUpdate")
		sd := s.stage_data.(*StageGame)
		if sd.ingame && sd.tid == 0 {
			// 发起任意请求：但测试中只发送获取玩家数据的请求
			sd.tid = s.timerMgr.CreateTimer(3*1000, false, nil)
		}
	}

	stages[idx].OnPacket = func(s *Session, packet *tcp.Packet) {
		if packet.Opcode == protocol.MSG_SC_ENTER_GAME {
			res := &msg.EnterGameResponse{}
			err := proto.Unmarshal(packet.Data, res)
			if err == nil {
				if res.ErrorCode == err_code.ERR_OK {
					sd := s.stage_data.(*StageGame)
					sd.ingame = true
					fmt.Println("MSG_ENTER_GAME Response")
				} else {
					fmt.Println("EnterGameResponse, err_code = ", res.ErrorCode)
				}
			} else {
				fmt.Println("proto.Unmarshal ERROR", err)
			}
		} else if packet.Opcode == protocol.MSG_SC_PlayerData {
			res := &msg.PlayerDataResponse{}
			err := proto.Unmarshal(packet.Data, res)
			if err == nil {
				fmt.Println("玩家数据：收到", res.Id, res.Acct, res.Name, res.Pid, res.Sid)
				sd := s.stage_data.(*StageGame)
				if res.Id != 0 && res.Id == sd.tid {
					sd.tid = 0
				}
			} else {
				fmt.Println("proto.Unmarshal ERROR", err)
			}
		}
	}

	stages[idx].OnTimer = func(s *Session, id uint64) {
		sd := s.stage_data.(*StageGame)
		if id == sd.tid {
			req := &msg.PlayerDataRequest{}
			req.Id = id
			s.SendPacket(protocol.MSG_CS_PlayerData, req)
			fmt.Println("玩家数据：请求中...")
		}
	}

}

func Next(s *Session) {
	stages[s.stage_id].OnLeave(s)
	s.stage_id++
	stages[s.stage_id].OnEnter(s)
}
