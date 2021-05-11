package referee

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/lobby"
	"gworld/ddz/lobby/smatch"
	"gworld/ddz/pb"
)

type handler = func(*Referee, comp.IMessage, comp.IMessage)
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
	_msg_executor[pb.Default_CreateMatchRequest_OP] = &executor_t{
		c: func() (comp.IMessage, comp.IMessage) { return &pb.CreateMatchRequest{}, &pb.CreateMatchResponse{} },
		h: handler_create_match,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_create_match(plr *Referee, req comp.IMessage, res comp.IMessage) {
	r := req.(*pb.CreateMatchRequest)
	s := req.(*pb.CreateMatchResponse)

	log.Info("create match: %v %v %v", r.TotalDeck, r.MatchName, r.Gamblers)

	if len(r.Gamblers) != 3 {
		s.ErrCode = gconst.Err_GamblerCount
		return
	}

	m := smatch.NewSMatch(&smatch.SMatchConf{
		TotalDeck:    r.TotalDeck,
		MatchMame:    r.MatchName,
		GamblerNames: r.Gamblers,
	})

	lobby.AddMatch(m)

	s.ErrCode = gconst.Err_OK
}
