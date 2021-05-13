package netmgr

import (
	"gworld/core/block"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/pb"
	"gworld/ddz_ai/args"
)

type handler = func(*connector, comp.IMessage)
type creator = func() comp.IMessage

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
	_msg_executor[pb.Default_RegisterResponse_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.RegisterResponse{} },
		h: handler_Register,
	}

	_msg_executor[pb.Default_SitResponse_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.SitResponse{} },
		h: handler_SitResponse,
	}

	_msg_executor[pb.Default_DealCardNotify_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.DealCardNotify{} },
		h: handler_DealCardNotify,
	}

	_msg_executor[pb.Default_CallScoreBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.CallScoreBroadcast{} },
		h: handler_CallScoreBroadcast,
	}

	_msg_executor[pb.Default_CallScoreResultBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.CallScoreResultBroadcast{} },
		h: handler_CallScoreResultBroadcast,
	}

	_msg_executor[pb.Default_CallScoreCalcBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.CallScoreCalcBroadcast{} },
		h: handler_CallScoreCalcBroadcast,
	}

	_msg_executor[pb.Default_PlayBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.PlayBroadcast{} },
		h: handler_PlayBroadcast,
	}

	_msg_executor[pb.Default_PlayResponse_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.PlayResponse{} },
		h: handler_PlayResponse,
	}

	_msg_executor[pb.Default_PlayResultBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.PlayResultBroadcast{} },
		h: handler_PlayResultBroadcast,
	}

	_msg_executor[pb.Default_DeckEndBroadcast_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.DeckEndBroadcast{} },
		h: handler_DeckEndBroadcast,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_Register(c *connector, res comp.IMessage) {
	s := res.(*pb.RegisterResponse)
	_ = s

	if s.ErrCode != gconst.Err_OK {
		block.Signal()
		return
	}

	c.SendMessage(&pb.SitRequest{
		MatchId: uint32(args.MatchID),
	})
}

func handler_SitResponse(c *connector, res comp.IMessage) {
	s := res.(*pb.SitResponse)

	if s.ErrCode != gconst.Err_OK {
		block.Signal()
		return
	}
}

func handler_DealCardNotify(c *connector, res comp.IMessage) {
	s := res.(*pb.DealCardNotify)
	ai.Init(c, s.Pos, s.Cards)
}

func handler_CallScoreBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.CallScoreBroadcast)

	ai.CallScoreBroadcast(s.Pos, s.History)
}

func handler_CallScoreResultBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.CallScoreResultBroadcast)

	ai.CallScoreResultBroadcast(s.Pos, s.Score)
}

func handler_CallScoreCalcBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.CallScoreCalcBroadcast)
	ai.CallScoreCalcBroadcast(s.Draw, s.Lord, s.Score, s.Cards)
}

func handler_PlayBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.PlayBroadcast)
	ai.PlayBroadcast(s.Pos, s.First)
}

func handler_PlayResponse(c *connector, res comp.IMessage) {
	s := res.(*pb.PlayResponse)
	ai.PlayResponse(s.ErrCode)
}

func handler_PlayResultBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.PlayResultBroadcast)
	ai.PlayResultBroadcast(s.Pos, s.Cards)
}

func handler_DeckEndBroadcast(c *connector, res comp.IMessage) {
	s := res.(*pb.DeckEndBroadcast)
	ai.DeckEndBroadcast(s.Score)
}
