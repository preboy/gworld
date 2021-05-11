package comp

// ----------------------------------------------------------------------------
// player mgr

var (
	GM IGamblerManager
	RM IRefereeManager
)

func Init_GamblerManager(mgr IGamblerManager) {
	GM = mgr
}

func Init_RefereeManager(mgr IRefereeManager) {
	RM = mgr
}

type IGamblerManager interface {
	NewGambler(pid string) IGambler
	FindGambler(pid string) IGambler
}

type IRefereeManager interface {
	NewReferee(pid string) IReferee
	FindReferee(pid string) IReferee
}
