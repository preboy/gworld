/*
   背景：
   在一个函数内可能会执行若干个模块的函数
   若其中一个模块出现panic,在有上层保护的情况下，虽不会导致程序崩溃，但后继的模块不会执行
   导致数据丢失等问题
*/

package utils

import (
	"gworld/core/log"
)

func ExecuteSafely(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[PANIC] STACK TRACE: %s", err)
			log.Error("%s", Callstack())
		}
	}()

	f()
}
