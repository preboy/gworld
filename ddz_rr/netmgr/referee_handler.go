package netmgr

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/pb"
	"strconv"

	"os/exec"
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
	_msg_executor[pb.Default_CreateMatchResponse_OP] = &executor_t{
		c: func() comp.IMessage { return &pb.CreateMatchResponse{} },
		h: handler_create_match,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_create_match(c *connector, res comp.IMessage) {
	s := res.(*pb.CreateMatchResponse)

	log.Info("create match response: %v", s.ErrCode)

	if s.ErrCode == gconst.Err_OK {
		for i := 0; i < len(s.Gamblers); i++ {
			cmd := exec.Command("ddz_ai.exe", strconv.Itoa(int(s.MatchID)), s.Gamblers[i])
			err := cmd.Run()
			if err != nil {
				log.Info("load ddz_ai failed: %v %v", s.MatchID, s.Gamblers[i])
			} else {
				log.Info("load ddz_ai ok: %v %v", s.MatchID, s.Gamblers[i])
			}
		}
	}
}
