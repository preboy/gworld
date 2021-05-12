package netmgr

import (
	"gworld/core/log"
	"gworld/ddz/comp"
	"gworld/ddz/gconst"
	"gworld/ddz/pb"
	"os"
	"strconv"

	"os/exec"
)

type handler = func(*connector, comp.IMessage, comp.IMessage)
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
	_msg_executor[pb.Default_CreateMatchRequest_OP] = &executor_t{
		c: func() (comp.IMessage, comp.IMessage) { return &pb.CreateMatchRequest{}, &pb.CreateMatchRequest{} },
		h: handler_create_match,
	}
}

// ----------------------------------------------------------------------------
// handlers

func handler_create_match(c *connector, req comp.IMessage, res comp.IMessage) {
	r := req.(*pb.CreateMatchRequest)
	s := req.(*pb.CreateMatchResponse)

	log.Info("create match response: %s %d", r.MatchName, s.ErrCode)

	if s.ErrCode == gconst.Err_OK {
		for i := 0; i < len(r.Gamblers); i++ {
			cmd := exec.Command("ddz_ai.exe", strconv.Itoa(int(s.MatchID)), r.Gamblers[i])
			err := cmd.Run()
			if err != nil {
				log.Info("load ddz_ai failed: %d %s", s.MatchID, r.Gamblers[i])
			} else {
				log.Info("load ddz_ai ok: %d %s", s.MatchID, r.Gamblers[i])
			}
		}
	}

	os.Exit(0)
}
