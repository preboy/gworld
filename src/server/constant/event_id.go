package constant

// EventID
const (
	EVT_APP_BEGIN uint32 = 0x100 + iota

	// player events
	EVT_PLR_LOGIN_FIRST
	EVT_PLR_LOGIN
	EVT_PLR_LOGOUT
	EVT_PLR_LEVEL_UP
	EVT_PLR_LEVEL_DEAD
)
