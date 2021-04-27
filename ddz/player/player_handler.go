package player

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby"
	"gworld/ddz/pb"
)

type handler = func(*Player, comp.Message, comp.Message)
type creator = func() (comp.Message, comp.Message)

var (
	_msg_executor = map[int32]*executor_t{}
)

type executor_t struct {
	c creator
	h handler
}

// ----------------------------------------------------------------------------
// init

func init() {
	// NOTE: 手写易出错, 此处注册的内容最好自动生成 (目前暂无此工具)

	_msg_executor[pb.Default_RegisterRequest_OP] = &executor_t{
		c: func() (comp.Message, comp.Message) { return &pb.RegisterRequest{}, &pb.RegisterResponse{} },
		h: handler_register,
	}

	_msg_executor[pb.Default_JoinRequest_OP] = &executor_t{
		c: func() (comp.Message, comp.Message) { return &pb.JoinRequest{}, &pb.JoinResponse{} },
		h: handler_join,
	}

	_msg_executor[pb.Default_CallScoreRequest_OP] = &executor_t{
		c: func() (comp.Message, comp.Message) { return &pb.CallScoreRequest{}, &pb.CallScoreResponse{} },
		h: handler_callscore,
	}

	_msg_executor[pb.Default_PlayRequest_OP] = &executor_t{
		c: func() (comp.Message, comp.Message) { return &pb.PlayRequest{}, &pb.PlayResponse{} },
		h: handler_play,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_register(plr *Player, req comp.Message, res comp.Message) {
	r := req.(*pb.RegisterRequest)
	s := req.(*pb.RegisterResponse)

	log.Info("login with: %s %s", r.Name, r.Pass)

	s.ErrCode = gconst.Err_OK
}

func handler_join(plr *Player, req comp.Message, res comp.Message) {
	// r := req.(*pb.JoinRequest)
	s := req.(*pb.JoinResponse)

	if lobby.Queue(plr.GetPID()) {
		s.ErrCode = gconst.Err_OK
	} else {
		s.ErrCode = gconst.Err_InLobbyOrMatch
	}
}

func handler_callscore(plr *Player, req comp.Message, res comp.Message) {
	lobby.OnMessage(plr.GetPID(), req, res)
}

func handler_play(plr *Player, req comp.Message, res comp.Message) {
	lobby.OnMessage(plr.GetPID(), req, res)
}
