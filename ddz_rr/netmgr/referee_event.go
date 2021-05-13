package netmgr

import (
	"gworld/core/log"
	"gworld/core/tcp"
	"gworld/core/utils"
	"gworld/ddz/pb"

	"github.com/gogo/protobuf/proto"
)

func referee_event_opened(c *connector) {
	msg := &pb.CreateMatchRequest{
		TotalDeck: 10,
		MatchName: "dev_test",
		Gamblers:  []string{"name_1", "name_2", "name_3"},
	}

	c.SendMessage(msg)

	log.Info("referee opened")
}

func referee_event_closed(c *connector) {
	log.Info("referee closed")
}

func referee_event_error(c *connector, err string) {
	log.Info("referee connect to ddz failed, err = %s", err)
}

func referee_handler(c *connector, p *tcp.Packet) {
	e, ok := _msg_executor[int32(p.Opcode)]
	if !ok {
		log.Warning("Unknown packet : %s %d", c.id, p.Opcode)
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
