package event

// EventID
const (
	EVT_SYS_BEGIN uint32 = 0 + iota

	// system event
	EVT_SYS_READY

	// sched events
	EVT_SCHED_MIN
	EVT_SCHED_HOUR
	EVT_SCHED_DAY
	EVT_SCHED_WEEK
	EVT_SCHED_MONTH
	EVT_SCHED_YEAR

	EVT_SCHED_SYNC_CALL
)
