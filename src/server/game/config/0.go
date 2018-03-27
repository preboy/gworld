package config

import (
	"core/log"
)

func Load() {
	log.Info("Loading Config Starting ...")

	load_global()
	load_creature()
	load_hero()
	load_aura()
	load_skill()
	load_item()
	load_creature_team()
	load_market_conf()

	log.Info("Loading Config COMPLETE !!!")
}
