package event

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
