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

	log.Info("Loading Config COMPLETE !!!")
}
