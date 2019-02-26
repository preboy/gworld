package constant

// EventID
const (
	EVT_APP_BEGIN uint32 = 0x100 + iota

	// ------------------------------------------------------------------------
	// system event

	EVT_SYS_ConfigLoaded
	EVT_SYS_SystemStart
	EVT_SYS_SystemReady

	// ------------------------------------------------------------------------
	// player events

	EVT_PLR_LOGIN_FIRST
	EVT_PLR_LOGIN
	EVT_PLR_LOGOUT
	EVT_PLR_LEVEL_UP
	EVT_PLR_KILL_MONSTER

	// ------------------------------------------------------------------------
	// hero events

	EVT_HERO_LEVEL_UP
)

const ()
