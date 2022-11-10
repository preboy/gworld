package netmgr

import (
	"gworld/ddz/loop"
)

var (
	_conn *connector
)

func init() {
	loop.Register(func() {
		update()
	})
}

// ----------------------------------------------------------------------------
// local

func update() {
	if _conn != nil {
		_conn.Update()
	}
}

// ----------------------------------------------------------------------------
// export

func Init() {
	_conn = NewConnector(ai_handler)
	_conn.Start("127.0.0.1:12345")

	_conn.On("error", func(c *connector, args []interface{}) {
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
		_conn = nil
	})
}

func Release() {
	if _conn != nil {
		_conn.Close()
	}
}
