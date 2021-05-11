package referee

import (
	"gworld/ddz/comp"
)

type handler = func(*Referee, comp.IMessage, comp.IMessage)
type creator = func() (comp.IMessage, comp.IMessage)

var (
	_msg_executor = map[int32]*executor_t{}
)

type executor_t struct {
	c creator
	h handler
}

// ----------------------------------------------------------------------------
// init

func init() {
	// NOTE: 手写易出错, 此处注册的内容最好自动生成 (目前暂无此工具)

}

// ----------------------------------------------------------------------------
// handlers
