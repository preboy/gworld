package args

import (
	"gworld/core/log"

	"os"
)

var (
	NickName  string
	MatchName string
)

func Parse() bool {
	if len(os.Args) != 3 {
		log.Info("ddz_ai <name> <match_name>")
		return false
	}

	NickName = os.Args[1]
	MatchName = os.Args[2]

	return true
}
