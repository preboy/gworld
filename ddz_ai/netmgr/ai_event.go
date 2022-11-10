package netmgr

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/pb"
	"gworld/ddz_ai/args"

	"github.com/gogo/protobuf/proto"
)

func ai_event_opened(c *connector) {
	log.Info("ai opened")

	c.SendMessage(&pb.RegisterRequest{
		Name: args.NickName,
	})

	_ai.SetConnector(c)
}

func ai_event_closed(c *connector) {
	log.Info("ai closed")

	_ai.SetConnector(nil)
}

func ai_event_error(c *connector, err string) {
	log.Info("ai connect to ddz failed, err = %s", err)
}

func ai_handler(c *connector, p *tcp.Packet) {
	e, ok := _msg_executor[int32(p.Opcode)]
	if !ok {
		log.Warning("Unknown packet : %v %v", c.id, p.Opcode)
		return
	}

	res := e.c()

	err := proto.Unmarshal(p.Data, res)
	if err != nil {
		log.Error("proto.Unmarshal ERROR: %v %v", c.id, p.Opcode)
		return
	}

	str := utils.ObjectToString(res)
	log.Info("RECV packet: %v, %v, %v", c.id, p.Opcode, str)

	e.h(c, res)
}
