package player

import (
	"fmt"
	"strings"

	"core/utils"
	"game/app"
	"game/constant"
)

// ----------------------------------------------------------------------------

func (self *Player) on_gm_command(args []string) int32 {
	fmt.Println("on_gm_command:", args)

	switch args[0] {

	case "save":
		self.Save()

	case "vip":
		if len(args) > 1 {
			val := utils.Atou32(args[1])
			self.data.Vip = val
			self.SendNotice("VipLevel: "+utils.U32toa(val), 0)
		}

	case "lv":
		if len(args) > 1 {
			val := utils.Atou32(args[1])
			self.data.Lv = val
			self.SendNotice("Level: "+utils.U32toa(val), 0)
		}

	case "item":
		for i := 1; i < len(args); i++ {
			ip := app.NewItemProxy(constant.ItemLog_GM)
			item := strings.Split(args[i], "|")
			id := utils.Atou32(item[0])
			ct := utils.Atou32(item[1])
			ip.Add(id, uint64(ct))
			ip.Apply(self)
			println("curr:", item[0], self.GetItem(id))
		}

	case "hero":
		for i := 1; i < len(args); i++ {
			item := strings.Split(args[i], "|")
			id := utils.Atou32(item[0])
			hero := self.GetHero(id)
			if hero == nil {
				self.AddHero(id)
			} else {
				self.SendNotice("Hero: "+item[0]+" already exist", 0)
			}
		}

	case "prop":
		str_props := "\n"
		for _, hero := range self.data.Heros {
			str := fmt.Sprintf("%d = %s", hero.Id, hero.ToBattleUnit().Prop.Dump())
			str_props += str
		}
		self.SendNotice(str_props, 0)

	default:
		println("unknown command:", args[0])
		return 0
	}
	return 1
}
