package event

// EventID
const (

	// system event
	EVT_SYSTEM_READY = iota

	// player evnet
	EVT_PLAYER_LOGIN
	EVT_PLAYER_LOGOUT
	EVT_PLAYER_LEVEL_UP
	EVT_PLAYER_LEVEL_DIED
)

type EventInfo struct {
	EvnetID int
}

type EventMgr struct {
	evts chan *EventInfo
	plr  IPlayerEventMgr
}

type IPlayerEventMgr interface {
	OnEvent(evt *EventInfo) int
}

func NewEventMgr(plr IPlayerEventMgr) *EventMgr {
	return &EventMgr{
		plr:  plr,
		evts: make(chan *EventInfo),
	}
}

func (self *EventMgr) Fire(evt *EventInfo) {
	self.evts <- evt
}

func (self *EventMgr) Update() {
	for {
		select {
		case evt := <-self.evts:
			self.plr.OnEvent(evt)
		default:
			break
		}
	}
}
