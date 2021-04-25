package comp

// ----------------------------------------------------------------------------
// player mgr

var (
	PM IPlayerManager
)

func Init_PlayerManager(mgr IPlayerManager) {
	PM = mgr
}

type IPlayerManager interface {
	FindPlayer(pid string) IPlayer
}
