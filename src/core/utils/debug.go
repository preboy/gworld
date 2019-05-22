package utils

import (
	"encoding/json"
	"runtime/debug"
)

func init() {
	// 这个开启之后，显示所有线程的堆栈
	// debug.SetTraceback("all")
}

func Callstack() string {
	return string(debug.Stack())
}

func JsonPretty(v interface{}) (ret string) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		ret = string(data)
	} else {
		ret = "JsonPretty Error !"
	}

	return
}
