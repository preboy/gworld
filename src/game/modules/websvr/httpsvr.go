package httpsvr

import (
	"net/http"

	"core/log"
)

// ============================================================================

func Start() {

	go func() {

		http.HandleFunc("/gm", dispatcher)

		if http.ListenAndServe(app.GetGameConfig().WebSvr, nil) != nil {
			log.Error("Start httpsvr FAILED:", err)
		}
	}()
}
