package event

// EventID
const (

	// system event
	EVT_SYS_READY uint32 = iota

	EVT_SCHED_MIN
	EVT_SCHED_HOUR
	EVT_SCHED_DAY
	EVT_SCHED_WEEK
	EVT_SCHED_MONTH
	EVT_SCHED_YEAR

	EVT_SCHED_SYNC_CALL

	// player evnet
	EVT_PLR_LOGIN_FIRST
	EVT_PLR_LOGIN
	EVT_PLR_LOGOUT

	EVT_PLR_LEVEL_UP
	EVT_PLR_LEVEL_DEAD
)

type Event struct {
	Id  uint32
	Ptr interface{}
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
		Id:  id,
		Ptr: ptr,
	}
}

func NewEventMgr(r IEventReceiver) *EventMgr {
	return &EventMgr{
		receiver: r,
		evts:     make(chan *Event, 0x100),
	}
}

func (self *EventMgr) Len() int {
	return len(self.evts)
}

func (self *EventMgr) Fire(evt *Event) {
	self.evts <- evt
}

func (self *EventMgr) Update() (busy bool) {
	for {
		select {
		case evt := <-self.evts:
			self.receiver.OnEvent(evt)
			busy = true
		default:
			return
		}
	}
}
