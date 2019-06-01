package websvr

import (
	"fmt"
	"net/http"

	"core/log"
	"game/loop"
)

var dispatcher = func(w http.ResponseWriter, r *http.Request) {
	var ret string
	var err error

	err = r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err2json(ret, err))
		log.Error("parsing data failed: %v", err)
		return
	}

	err = ErrNoKey
	chn := make(chan string)

	key := r.FormValue("key")
	fun := handlers[key]

	if fun != nil {

		loop.Get().PostFunc(func() {
			ret, err = fun(r)
			chn <- ret
			close(chn)
		})

		ret = <-chn
	}

	fmt.Fprint(w, err2json(ret, err))
}
