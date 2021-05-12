package args

import (
	"gworld/core/log"

	"os"
	"strconv"
)

var (
	MatchID  int
	NickName string
)

func Parse() bool {
	if len(os.Args) != 3 {
		log.Info("ddz_ai <match_id> <name>")
		return false
	}

	var err error
	MatchID, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Info("invalid match_id: %v", os.Args[1])
		return false
	}

	NickName = os.Args[2]
	if err != nil {
		log.Info("invalid match_id: %v", os.Args[2])
		return false
	}

	return true
}
