package websvr

import (
	"fmt"
	"net/http"

	"core"
	"core/logger"
	"game/app"
)

// ============================================================================

func r2json(r string, err error) string {
	if err == nil {
		if r == "" {
			return `{"err": ""}`
		} else {
			return r
		}
	} else {
		return fmt.Sprintf(`{"err": %q}`, err)
	}
}

func get_player(req *http.Request) (plr *app.Player, err error) {
	id := get_string(req, "plrid")

	plr = app.PlayerMgr.LoadPlayer(id, true)
	if plr == nil {
		err = ErrNoPlayer
	}

	return
}

func get_guild(req *http.Request) (gld *guild.Guild, err error) {
	id := get_string(req, "gldid")
	name := get_string(req, "gldname")

	if id != "" {
		gld = guild.GuildMgr.FindGuild(id)
	} else if name != "" {
		gld = guild.GuildMgr.FindGuildByName(name)
	}

	if gld == nil {
		err = ErrNoGuild
	}

	return
}

func get_int32(req *http.Request, k string) int32 {
	return core.Atoi32(req.FormValue(k))
}

func get_int64(req *http.Request, k string) int64 {
	return core.Atoi64(req.FormValue(k))
}

func get_string(req *http.Request, k string) string {
	return req.FormValue(k)
}

func get_int32_arr(req *http.Request, k string) (arr []int32) {
	for _, v := range req.Form[k] {
		arr = append(arr, core.Atoi32(v))
	}

	return
}

func get_int64_arr(req *http.Request, k string) (arr []int64) {
	for _, v := range req.Form[k] {
		arr = append(arr, core.Atoi64(v))
	}

	return
}

func get_string_arr(req *http.Request, k string) (arr []string) {
	for _, v := range req.Form[k] {
		arr = append(arr, v)
	}

	return
}
