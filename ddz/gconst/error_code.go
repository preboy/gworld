package gconst

const (
	// OK
	Err_Error = 0
	Err_OK    = 1

	// login
	Err_InvalidLogin = iota + 100

	// ddz logic
	Err_InLobbyOrMatch = iota + 200

	// sit
	Err_MatchNotExist

	// call
	Err_CallPid
	Err_CallPos
	Err_CallScore
	Err_CallScore2

	// play
	Err_NotYourTurn
	Err_CardNull
	Err_CardInvalid
	Err_CardNotExist
	Err_CardTypeInvalid
)

const (
	Err_GamblerCount = iota + 700
)
