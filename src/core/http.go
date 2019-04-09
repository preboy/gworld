package core

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ============================================================================

func HttpGet(addr string) (ret string) {
	res, err := http.Get(addr)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return string(body)
}

func HttpPost(addr string, data url.Values) (ret string) {
	res, err := http.PostForm(addr, data)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return string(body)
}

func HttpPostJson(addr string, J string) (ret string) {
	res, err := http.Post(addr, "application/json", bytes.NewBufferString(J))
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return string(body)
}
