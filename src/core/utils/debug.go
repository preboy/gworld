package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

func init() {
	// 这个开启之后，显示所有线程的堆栈
	debug.SetTraceback("all")
}

func Callstack() string {
	return string(debug.Stack())
}

func PrintPretty(v interface{}, mark string) {
	fmt.Printf("====== [%s] ====== (at:%s)\n", mark, time.Now())
	data, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		data = append(data, '\n')
		os.Stdout.Write(data)
	} else {
		fmt.Println("PrintPretty Error:", err)
	}
}
