package log

import (
	"fmt"
	"os"
)

type Logger struct {
	f      *os.File
	q      chan string
	quit   chan bool
	screen bool
}

// public
func (self *Logger) Info(format string, a ...interface{}) {
	s := "[" + get_time_string() + "] " + fmt.Sprintf(format, a...) + "\n"
	self._write_string(s)
}

func (self *Logger) Debug(format string, a ...interface{}) {
	s := "\033[1;36m [" + get_time_string() + "] " + fmt.Sprintf(format, a...) + "\033[m\n"
	self._write_string(s)
}

func (self *Logger) Warning(format string, a ...interface{}) {
	s := "\033[1;33m [" + get_time_string() + "] " + fmt.Sprintf(format, a...) + "\033[m\n"
	self._write_string(s)
}

func (self *Logger) Error(format string, a ...interface{}) {
	s := "\033[1;31m [" + get_time_string() + "] " + fmt.Sprintf(format, a...) + "\033[m\n"
	self._write_string(s)
}

func (self *Logger) Fatal(format string, a ...interface{}) {
	s := "\033[1;35m [" + get_time_string() + "] " + fmt.Sprintf(format, a...) + "\033[m\n"
	self._write_string(s)
}

func (self *Logger) Go() {
	go func() {
		for {
			select {
			case s := <-self.q:
				self.f.WriteString(s)
				if self.screen {
					fmt.Println(s)
				}
			case <-self.quit:
				break
			}

		}
		self.f.Close()
	}()
}

func (self *Logger) Stop() {
	self.quit <- true
}

// private
func (self *Logger) _write_string(s string) {
	self.q <- s
}

// public module function
func NewLogger(name string, screen bool) *Logger {
	f, err := os.Create(name)
	if err != nil {
		return nil
	}

	l := &Logger{
		f:      f,
		q:      make(chan string),
		quit:   make(chan bool),
		screen: screen,
	}
	l.Go()
	return l
}
