package utils

import (
	"gworld/core/log"

	"encoding/json"
	"io/ioutil"
)

var (
	LoadJsonAsObj = load_from_json
	LoadJsonAsArr = load_from_json
)

func load_from_json(file string, arr interface{}) bool {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("loading [%s] failed! err = %s", file, err)
		return false
	}

	err = json.Unmarshal(content, &arr)
	if err != nil {
		log.Error("Unmarshal [%s] failed! %s", file, err)
		return false
	}

	return true
}
