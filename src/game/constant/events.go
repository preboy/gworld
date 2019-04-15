package constant

// EventID
const (
	EVT_APP_BEGIN uint32 = 0x100 + iota

	// ------------------------------------------------------------------------
	// system event

	EVT_SYS_ConfigLoaded
	EVT_SYS_SystemStart
	EVT_SYS_SystemReady

	Evt_Auth // 用户名验证成功

	// ------------------------------------------------------------------------
	// player events

	Evt_Plr_LoginFirst
	Evt_Plr_Login
	Evt_Plr_Logout
	Evt_Plr_LevelUp
	Evt_Plr_KillMonster
	Evt_Plr_DataReset // 跨天重置用户数据

	// ------------------------------------------------------------------------
	// hero events

	EVT_HERO_LEVEL_UP
)

const ()
