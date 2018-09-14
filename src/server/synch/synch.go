package synch

import (
	"sync"

	"core/event"
)

type ISync interface {
	FireEvent(evt *event.Event)
}

// 将函数在一个ISync(通常他有一个routine且能处理事件，例如：player   main_loop)中执行
// NOTE：为防止嵌套:  1. 请谨慎使用本函数	2. f应尽量简单 3. 不能往自己身上发
func SyncExecute(routine ISync, f func()) {
	w := &sync.WaitGroup{}
	w.Add(1)

	_f := func() {
		defer w.Done()
		f()
	}

	evt := event.NewEvent(event.EVT_SCHED_SYNC_CALL, _f)
	routine.FireEvent(evt)

	w.Wait()
}
