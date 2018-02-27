package schedule

import (
	"core/event"
	"core/thread"
	"time"
)

var _curr_date *date

type date struct {
	_curr_year  int
	_curr_month int
	_curr_week  int
	_curr_day   int
	_curr_hour  int
	_curr_min   int
}

type ISchedule interface {
	OnSchedule(evt *event.Event)
}

var (
	_observer map[string]ISchedule
	_thread   *thread.Thread
)

func init() {
	_observer = make(map[string]ISchedule)
}

func new_date() *date {
	now := time.Now()
	return &date{
		_curr_year:  now.Year(),
		_curr_month: int(now.Month()),
		_curr_week:  int(now.Weekday()),
		_curr_day:   now.Day(),
		_curr_hour:  now.Hour(),
		_curr_min:   now.Minute(),
	}
}

func update_time() {
	date := new_date()

	if _curr_date._curr_hour == date._curr_hour {
		return
	}

	// new hour
	if _curr_date._curr_hour != date._curr_hour {
		evt := event.NewEvent(event.EVT_SCHED_HOUR, nil)
		for _, obj := range _observer {
			obj.OnSchedule(evt)
		}
	}

	// new day
	if _curr_date._curr_day != date._curr_day {
		evt := event.NewEvent(event.EVT_SCHED_DAY, nil)
		for _, obj := range _observer {
			obj.OnSchedule(evt)
		}
	}

	// new week
	if _curr_date._curr_week != date._curr_week {
		evt := event.NewEvent(event.EVT_SCHED_WEEK, nil)
		for _, obj := range _observer {
			obj.OnSchedule(evt)
		}
	}

	// new month
	if _curr_date._curr_month != date._curr_month {
		evt := event.NewEvent(event.EVT_SCHED_MONTH, nil)
		for _, obj := range _observer {
			obj.OnSchedule(evt)
		}
	}

	// new year
	if _curr_date._curr_year != date._curr_year {
		evt := event.NewEvent(event.EVT_SCHED_YEAR, nil)
		for _, obj := range _observer {
			obj.OnSchedule(evt)
		}
	}

	_curr_date = date
}

func Start() {
	_curr_date = new_date()
	if _thread == nil {
		_thread = thread.NewThread(update_time, 1000)
		_thread.Go()
	}
}

func Stop() {
	if _thread != nil {
		_thread.Stop()
	}
}

func Register(name string, obj ISchedule) bool {
	if Exist(name) != nil {
		return false
	}
	_observer[name] = obj
	return true
}

func UnRegister(name string) {
	delete(_observer, name)
}

func Exist(name string) ISchedule {
	if obj, ok := _observer[name]; ok {
		return obj
	}
	return nil
}

func Year() int {
	return _curr_date._curr_year
}

func Month() int {
	return _curr_date._curr_month
}

func Week() int {
	return _curr_date._curr_week
}

func Day() int {
	return _curr_date._curr_day
}

func Hour() int {
	return _curr_date._curr_hour
}

func Min() int {
	return _curr_date._curr_min
}
