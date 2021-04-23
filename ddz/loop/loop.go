package loop

import (
	"gworld/core/log"
	"time"
)

// ----------------------------------------------------------------------------
// local
func update() {
	log.Info("update")
}

// ----------------------------------------------------------------------------
// export

func Run() {
	go func() {
		for {
			update()
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
