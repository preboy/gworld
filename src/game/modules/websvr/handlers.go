package websvr

import (
	"errors"
	"net/http"

	"core/utils"
)

// ============================================================================

var (
	ErrNoKey    = errors.New("invalid key")
	ErrArgs     = errors.New("invalid params")
	ErrNoPlayer = errors.New("player not found")
	ErrNoGuild  = errors.New("guild not found")
	ErrNoHero   = errors.New("hero not found")
	ErrNoArmor  = errors.New("armor not found")
	ErrNoConf   = errors.New("conf file not found")
)

// ============================================================================

var handlers = map[string]func(*http.Request) (string, error){
	"plrinfo": handle_plrinfo,
}

// ============================================================================

func handle_plrinfo(req *http.Request) (ret string, err error) {
	plr, err := get_player()
	if err != nil {
		return
	}

	ret = utils.ObjectToString(plr)

	return
}


