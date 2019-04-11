package microsvr

import (
	"core/log"
)

func Start() {
	var names = map[string]*Svr{}

	for _, v := range svrs {
		key := v.mod.GetName()
		if _, ok := names[key]; ok {
			log.Error("Duplicate Name for '%s'", key)
			continue
		}
		names[key] = v

		v.Start()
	}
}

func Stop() {
	for _, v := range svrs {
		v.Stop()
	}
}
