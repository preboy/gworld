package event

// EventID
const (

	// system event
	EVT_SYS_READY uint32 = iota

	// player evnet
	EVT_PLR_LOGIN
	EVT_PLR_LOGOUT
	EVT_PLR_LEVEL_UP
	EVT_PLR_LEVEL_DEAD
)

type Event struct {
	id  uint32
	ptr interface{}
}

type EventMgr struct {
	evts     chan *Event
	receiver IEventReceiver
}

type IEventReceiver interface {
	OnEvent(evt *Event)
}

func NewEvent(id uint32, ptr interface{}) *Event {
	return &Event{
		id:  id,
		ptr: ptr,
	}
}

func NewEventMgr(r IEventReceiver) *EventMgr {
	return &EventMgr{
		receiver: r,
		evts:     make(chan *Event, 0x100),
	}
}

func (self *EventMgr) Fire(evt *Event) {
	self.evts <- evt
}

func (self *EventMgr) Update() {
	for {
		select {
		case evt := <-self.evts:
			self.receiver.OnEvent(evt)
		default:
			break
		}
	}
}
