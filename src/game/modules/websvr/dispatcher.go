package websvr

import (
	"fmt"
	"net/http"

	"game/loop"
)

var dispatcher = func(w http.ResponseWriter, req *http.Request) {
	var r string
	var err error

	err = req.ParseForm()
	if err != nil {
		fmt.Fprint(w, r2json(r, err))
		log.Error("GM request data parsing failed:", err)
		return
	}

	ch := make(chan string)
	err = ErrNoKey

	key := req.FormValue("key")
	h := handlers[key]
	if h != nil {
		loop.Push(func() {
			r, err = h(req)
			ch <- r
			close(ch)
		})
		r = <-ch
	}

	fmt.Fprint(w, r2json(r, err))
}
