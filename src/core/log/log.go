package log

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	w      *sync.WaitGroup
	f      *os.File
	q      chan string
	quit   chan bool
	screen bool
}

// public
func (self *Logger) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (self *Logger) Info(format string, a ...interface{}) {
	s := log_prefix(1) + fmt.Sprintf(format, a...) + log_suffix()
	self._write_string(s)
}

func (self *Logger) Debug(format string, a ...interface{}) {
	s := log_prefix(2) + fmt.Sprintf(format, a...) + log_suffix()
	self._write_string(s)
}

func (self *Logger) Warning(format string, a ...interface{}) {
	s := log_prefix(3) + fmt.Sprintf(format, a...) + log_suffix()
	self._write_string(s)
}

func (self *Logger) Error(format string, a ...interface{}) {
	s := log_prefix(4) + fmt.Sprintf(format, a...) + log_suffix()
	self._write_string(s)
}

func (self *Logger) Fatal(format string, a ...interface{}) {
	s := log_prefix(5) + fmt.Sprintf(format, a...) + log_suffix()
	self._write_string(s)
}

func (self *Logger) Go() {
	go func() {
		defer func() {
		E:
			for {
				select {
				case s := <-self.q:
					self.f.WriteString(s + "\n")
				default:
					break E
				}
			}
			self.f.Close()
			self.w.Done()
		}()

		self.w.Add(1)

		for {
			select {
			case s := <-self.q:
				self.f.WriteString(s + "\n")
				if self.screen {
					fmt.Println(s)
				}
			case <-self.quit:
				return
			}
		}
	}()
}

func (self *Logger) Stop() {
	self.quit <- true
	self.w.Wait()
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
		w:      &sync.WaitGroup{},
		f:      f,
		q:      make(chan string, 0x100),
		quit:   make(chan bool),
		screen: screen,
	}
	l.Go()
	return l
}
