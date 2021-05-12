package netmgr

import (
	"gworld/ddz/loop"
)

var (
	_referee *connector
)

func init() {
	loop.Register(func() {
		update()
	})
}

// ----------------------------------------------------------------------------
// local

func update() {
	if _referee != nil {
		_referee.Update()
	}
}

// ----------------------------------------------------------------------------
// export

func Init() {
	_referee = NewConnector(referee_handler)
	_referee.Start("127.0.0.1:12346")

	_referee.On("error", func(c *connector, args []interface{}) {
		loop.Post(func() {
			referee_event_error(c, args[0].(string))
		})
	}).On("opened", func(c *connector, args []interface{}) {
		loop.Post(func() {
			referee_event_opened(c)
		})
	}).On("closed", func(c *connector, args []interface{}) {
		loop.Post(func() {
			referee_event_closed(c)
		})
		_referee = nil
	})
}

func Release() {
	if _referee != nil {
		_referee.Close()
	}
}
