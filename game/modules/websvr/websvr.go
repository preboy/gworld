package websvr

import (
	"net/http"

	"gworld/core/log"
	"gworld/game/app"
)

// ============================================================================

func Start() {

	go func() {

		http.HandleFunc("/gm", dispatcher)

		if err := http.ListenAndServe(app.GetGameConfig().WebSvr, nil); err != nil {
			log.Error("Start websvr FAILED:", err)
		}
	}()
}
