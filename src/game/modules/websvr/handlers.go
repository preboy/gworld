package websvr

import (
	"errors"
	"net/http"
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

var handlers = map[string]func(*http.Request) (r string, err error){
	"plrinfo": handle_plrinfo,
}

// ============================================================================

func handle_plrinfo(req *http.Request) (r string, err error) {
	r = "this is plr info ..."
	return
}
