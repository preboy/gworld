package clients

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Player struct {
	pid    uint32
	sid    uint32
	name   string
	socket *net.TCPConn
	data   uint32
	// 再搞几个通道用来通信
	ch_packet chan Packet
}

var _plrs_map = make(map[uint32]*Player)
var _plrs_arr = make([]*Player, 0x1000)

var _index uint32

func CreatePlayer(name string) *Player {

	var plr = new(Player)
	plr.pid = rand.Uint32()
	plr.sid = _index
	plr.name = name
	plr.socket = nil
	plr.ch_packet = make(chan Packet, 16)

	_plrs_arr[plr.sid] = plr
	_plrs_map[plr.pid] = plr

	_index++

	go player_loop(plr)

	return plr
}

func GetPlayerByIndex(sid uint32) *Player {
	if _index < 0x1000 {
		return _plrs_arr[sid]
	}
	return nil
}

func GetPlayerById(pid uint32) *Player {
	plr, ok := _plrs_map[pid]
	if !ok {
		return nil
	}
	return plr
}

func player_loop(plr *Player) {

	for {
		select {
		case pck := <-plr.ch_packet:
			fmt.Println("new packet", pck.code, len(pck.data))

		// case <-time.After(1 * time.Minute):
		// fmt.Println("1 Minute later")
		default:
			// fmt.Println("not get packet")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
