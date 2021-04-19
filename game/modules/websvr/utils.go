package websvr

import (
	"fmt"
	"net/http"

	"gworld/core/utils"
	"gworld/game/player"
)

// ============================================================================

func err2json(ret string, err error) string {
	if err == nil {
		if ret == "" {
			return `{"err": ""}`
		} else {
			return ret
		}
	} else {
		return fmt.Sprintf(`{"err": %q}`, err)
	}
}

// ============================================================================

func get_int32(req *http.Request, k string) int32 {
	return utils.Atoi32(req.FormValue(k))
}

func get_int64(req *http.Request, k string) int64 {
	return utils.Atoi64(req.FormValue(k))
}

func get_string(req *http.Request, k string) string {
	return req.FormValue(k)
}

func get_int32_arr(req *http.Request, k string) (arr []int32) {
	for _, v := range req.Form[k] {
		arr = append(arr, utils.Atoi32(v))
	}

	return
}

func get_int64_arr(req *http.Request, k string) (arr []int64) {
	for _, v := range req.Form[k] {
		arr = append(arr, utils.Atoi64(v))
	}

	return
}

func get_string_arr(req *http.Request, k string) (arr []string) {
	for _, v := range req.Form[k] {
		arr = append(arr, v)
	}

	return
}

// ============================================================================

func get_player(req *http.Request) (plr *player.Player, err error) {
	pid := get_string(req, "pid")

	plr = player.GetPlayer(pid)
	if plr == nil {
		err = ErrNoPlayer
	}

	return
}
