package gambler

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby"
	"gworld/ddz/pb"
)

type handler = func(*Gambler, comp.IMessage, comp.IMessage)
type creator = func() (comp.IMessage, comp.IMessage)

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
		c: func() (comp.IMessage, comp.IMessage) { return &pb.RegisterRequest{}, &pb.RegisterResponse{} },
		h: handler_register,
	}

	_msg_executor[pb.Default_SitRequest_OP] = &executor_t{
		c: func() (comp.IMessage, comp.IMessage) { return &pb.SitRequest{}, &pb.SitResponse{} },
		h: handler_sit,
	}

	_msg_executor[pb.Default_CallScoreRequest_OP] = &executor_t{
		c: func() (comp.IMessage, comp.IMessage) { return &pb.CallScoreRequest{}, &pb.CallScoreResponse{} },
		h: handler_callscore,
	}

	_msg_executor[pb.Default_PlayRequest_OP] = &executor_t{
		c: func() (comp.IMessage, comp.IMessage) { return &pb.PlayRequest{}, &pb.PlayResponse{} },
		h: handler_play,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_register(plr *Gambler, req comp.IMessage, res comp.IMessage) {
	r := req.(*pb.RegisterRequest)
	s := req.(*pb.RegisterResponse)

	log.Info("register with: %s", r.Name)

	if comp.GM.ExistGambler(r.Name) {
		s.ErrCode = gconst.Err_Error
	} else {
		s.ErrCode = gconst.Err_OK
		plr.Name = r.Name
	}
}

func handler_sit(plr *Gambler, req comp.IMessage, res comp.IMessage) {
	r := req.(*pb.SitRequest)
	s := req.(*pb.SitResponse)

	m := lobby.GetMatch(r.MatchId)
	if m == nil {
		s.ErrCode = gconst.Err_MatchNotExist
		return
	}

	if m.Sit(plr.GetPID()) {
		plr.Data.MatchID = r.MatchId
	}
}

func handler_callscore(plr *Gambler, req comp.IMessage, res comp.IMessage) {
	mid := plr.Data.MatchID
	if mid != 0 {
		m := lobby.GetMatch(mid)
		if m != nil {
			m.OnMessage(plr.GetPID(), req, res)
		}
	}
}

func handler_play(plr *Gambler, req comp.IMessage, res comp.IMessage) {
	mid := plr.Data.MatchID
	if mid != 0 {
		m := lobby.GetMatch(mid)
		if m != nil {
			m.OnMessage(plr.GetPID(), req, res)
		}
	}
}
