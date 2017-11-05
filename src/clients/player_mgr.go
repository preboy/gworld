package clients

import (
	_ "fmt"
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
	evtMgr    *EventMgr
}

var _plrs_map = make(map[uint32]*Player)
var _plrs_arr = make([]*Player, 0x1000)

var _index uint32
var last_update int64

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

func prepare_update(plr *Player) bool {
	now := time.Now().Unix()
	if now-100 >= last_update {
		last_update = now
		plr.OnUpdate()
		return true
	}
	return false
}

func player_loop(plr *Player) {
	for {
		select {
		case packet := <-plr.ch_packet:
			Dispatcher(&packet, plr)
		default:
			if !prepare_update(plr) {
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}

// ----------------- player evnet -----------------

func (plr *Player) OnUpdate() {
	plr.evtMgr.Loop()
}

func (plr *Player) OnEvent(evt *EventInfo) int {
	return 0
}
