package player

import (
	"gworld/core/log"
	"gworld/ddz/pb"

	"github.com/gogo/protobuf/proto"
)

type handler = func(proto.Message, proto.Message)
type creator = func() (proto.Message, proto.Message)

var (
	_msg_executor = map[int]*executor_t{}
)

type executor_t struct {
	c creator
	h handler
}

// ----------------------------------------------------------------------------
// init

func init() {

	_msg_executor[int(pb.RegisterRequest_OpCode)] = &executor_t{
		c: func() (proto.Message, proto.Message) { return &pb.RegisterRequest{}, &pb.RegisterResponse{} },
		h: handler_register,
	}

}

// ----------------------------------------------------------------------------
// handlers

func handler_register(req proto.Message, res proto.Message) {
	r := req.(*pb.RegisterRequest)
	s := req.(*pb.RegisterResponse)

	log.Info("%s", r.Name)
	log.Info("%d", s.Ret)
}
