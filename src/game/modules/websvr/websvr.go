package websvr

import (
	"net/http"

	"core/log"
	"game/app"
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
