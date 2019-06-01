package websvr

import (
	"core"
	"encoding/json"
	"fmt"
	"game/app"
	"game/app/dbmgr"
	"game/app/gamedata"
	"game/app/modules/guild"
	"sort"
	"strings"
	"time"
)

// ============================================================================

func get_plrinfo_base(plr *app.Player) (r string, err error) {
	user := plr.User()

	var tab [][]interface{}

	tab = append(tab, []interface{}{"Num", "Field", "Value"})

	tab = append(tab, []interface{}{"01", "UserId", user.Id})
	tab = append(tab, []interface{}{"02", "AccountId", user.IdDefault})
	tab = append(tab, []interface{}{"03", "Svr0", user.Svr0})
	tab = append(tab, []interface{}{"04", "Svr", user.Svr})
	tab = append(tab, []interface{}{"05", "Sdk", user.Sdk})
	tab = append(tab, []interface{}{"06", "Channel", user.Chn})
	tab = append(tab, []interface{}{"07", "Plat", user.Plat})
	tab = append(tab, []interface{}{"08", "DeviceId", user.DevId})
	tab = append(tab, []interface{}{"09", "Name", user.Name})
	tab = append(tab, []interface{}{"10", "Level", user.Lv})
	tab = append(tab, []interface{}{"11", "VIP", user.Vip.Level})
	tab = append(tab, []interface{}{"12", "TotalBill", user.Bill.GetSumString()})
	tab = append(tab, []interface{}{"13", "GhostId", user.Ghost.Id})
	tab = append(tab, []interface{}{"14", "WLevelRatio", fmt.Sprintf("%d%%", user.WLevel.GetExplorePct())})
	tab = append(tab, []interface{}{"15", "Guild", plr.GetGuildName()})
	tab = append(tab, []interface{}{"16", "CreateTime", user.CreateTs})
	tab = append(tab, []interface{}{"17", "LoginTime", user.LoginTs})
	tab = append(tab, []interface{}{"18", "LoginIP", user.LoginIP})

	// 在线情况
	{
		var text string

		if plr.IsOnline() {
			text = "<b style='color:green'>Online</b>"
		} else {
			text = "Offline"
		}

		tab = append(tab, []interface{}{"19", "Online-Status", text})
	}

	// 封号情况
	{
		var ban_text string
		ban_ts := dbmgr.Center_GetBanInfo(user.Id)

		if ban_ts.After(time.Now()) {
			ban_text = "<b style='color:red'>Banned</b>"
		} else {
			ban_text = "Normal"
		}

		tab = append(tab, []interface{}{"20", "Ban-Status", ban_text})
		tab = append(tab, []interface{}{"21", "BanUtil", ban_ts})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_plrinfo_ccy(plr *app.Player) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Name", "Amount"})

	for k, v := range plr.GetBag().Ccy {
		name := core.I32toa(k)
		conf := gamedata.ConfCurrency.Query(k)
		if conf != nil {
			name = conf.Name
		}

		tab = append(tab, []interface{}{name, v})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_plrinfo_hero(plr *app.Player) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Name", "Level", "Star", "Break", "Wake", "Equippment"})

	for _, hero := range plr.GetBag().Heroes {
		name := core.I32toa(hero.Id)
		conf := gamedata.ConfHero.Query(hero.Id)
		if conf != nil {
			name = conf.HeroName
		}

		var ar_buf []string
		for _, ar := range hero.ArmorBox() {
			if ar == nil {
				continue
			}

			name := core.I32toa(ar.Id)
			conf := gamedata.ConfItem.Query(ar.Id)
			if conf != nil {
				name = conf.Alias
			}

			ar_buf = append(ar_buf, fmt.Sprintf("%s:%d", name, ar.Lv))
		}

		tab = append(tab, []interface{}{name, hero.Lv, hero.Star, hero.Cls, hero.AwkLv, strings.Join(ar_buf, "<br>")})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_plrinfo_item(plr *app.Player) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Name", "Amount"})

	for k, v := range plr.GetBag().Items {
		name := core.I32toa(k)
		conf := gamedata.ConfItem.Query(k)
		if conf != nil {
			name = conf.Alias
		}

		tab = append(tab, []interface{}{name, v})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_plrinfo_armor(plr *app.Player) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Seq", "Name", "Level"})

	for _, ar := range plr.GetBag().Armors {
		name := core.I32toa(ar.Id)
		conf := gamedata.ConfItem.Query(ar.Id)
		if conf != nil {
			name = conf.Alias
		}

		tab = append(tab, []interface{}{ar.Seq, name, ar.Lv})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_plrinfo_guild(plr *app.Player) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Num", "Field", "Value"})

	gld := plr.GetGuild()
	if gld != nil {
		tab = append(tab, []interface{}{"01", "Id", gld.GetId()})
		tab = append(tab, []interface{}{"02", "Name", gld.GetName()})
		tab = append(tab, []interface{}{"03", "Owner", gld.Owner().GetName()})
		tab = append(tab, []interface{}{"04", "Level", gld.GetLevel()})
		tab = append(tab, []interface{}{"05", "MemberCount", len(gld.Members)})
		tab = append(tab, []interface{}{"06", "Rank", plr.GetGuildRank()})
	}

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

// ============================================================================

func get_gldinfo_base(gld *guild.Guild) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Num", "Field", "Value"})

	tab = append(tab, []interface{}{"01", "Id", gld.GetId()})
	tab = append(tab, []interface{}{"02", "Name", gld.GetName()})
	tab = append(tab, []interface{}{"03", "Owner", gld.Owner().GetName()})
	tab = append(tab, []interface{}{"04", "Level", gld.GetLevel()})
	tab = append(tab, []interface{}{"05", "MemberCount", len(gld.Members)})
	tab = append(tab, []interface{}{"06", "AtkPwr", gld.GetAtkPower()})
	tab = append(tab, []interface{}{"07", "Rank", gld.GetRank()})

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}

func get_gldinfo_members(gld *guild.Guild) (r string, err error) {
	var tab [][]interface{}

	tab = append(tab, []interface{}{"Id", "Name", "Level", "AtkPwr", "Rank", "Donation", "WX", "Online"})

	rk_dict := map[int32]string{
		1: "<b style='color:#ff6000'>Owner</b>",
		2: "VP",
		3: "Member",
	}

	ol_dict := map[bool]string{
		true:  "<b style='color:green'>Y</b>",
		false: "--",
	}

	for _, mb := range gld.Members {
		plr := app.PlayerMgr.LoadPlayer(mb.Id, true)
		if plr == nil {
			return
		}

		tab = append(tab, []interface{}{
			plr.GetId(),
			plr.GetName(),
			plr.GetLevel(),
			plr.GetAtkPower(),
			rk_dict[mb.Rank],
			mb.Donation,
			mb.Wuxun,
			ol_dict[plr.IsOnline()],
		})
	}

	sort.Slice(tab[1:], func(i, j int) bool {
		return tab[i+1][3].(int32) > tab[j+1][3].(int32)
	})

	// marshal
	data, err := json.Marshal(&tab)
	if err == nil {
		r = string(data)
	}

	return
}
