package netmgr

import (
	"gworld/ddz/loop"
)

var (
	_ai *connector
)

func init() {
	loop.Register(func() {
		update()
	})
}

// ----------------------------------------------------------------------------
// local

func update() {
	if _ai != nil {
		_ai.Update()
	}
}

// ----------------------------------------------------------------------------
// export

func Init() {
	_ai = NewConnector(ai_handler)
	_ai.Start("127.0.0.1:12345")

	_ai.On("error", func(c *connector, args []interface{}) {
		loop.Post(func() {
			ai_event_error(c, args[0].(string))
		})
	}).On("opened", func(c *connector, args []interface{}) {
		loop.Post(func() {
			ai_event_opened(c)
		})
	}).On("closed", func(c *connector, args []interface{}) {
		loop.Post(func() {
			ai_event_closed(c)
		})
		_ai = nil
	})
}

func Release() {
	if _ai != nil {
		_ai.Close()
	}
}
