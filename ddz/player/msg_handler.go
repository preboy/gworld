package player

import (
	"gworld/core/log"
	"gworld/ddz/pb"

	"github.com/gogo/protobuf/proto"
)

type Message interface {
	proto.Message
	GetOP() int32
}

type handler = func(Message, Message)
type creator = func() (Message, Message)

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
		c: func() (Message, Message) { return &pb.RegisterRequest{}, &pb.RegisterResponse{} },
		h: handler_register,
	}

}

// ----------------------------------------------------------------------------
// handlers

func handler_register(req Message, res Message) {
	r := req.(*pb.RegisterRequest)
	s := req.(*pb.RegisterResponse)

	log.Info("%s", r.Name)
	log.Info("%d", s.Ret)
}
