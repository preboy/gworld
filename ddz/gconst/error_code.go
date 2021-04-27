package gconst

const (
	// OK
	Err_Error = 0
	Err_OK    = 1

	// login
	Err_InvalidLogin = iota + 100

	// ddz logic
	Err_InLobbyOrMatch = iota + 200
)
