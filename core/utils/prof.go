package utils

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	_fcpu *os.File
)

func StartPprof(filename string) bool {
	if _fcpu != nil {
		fmt.Println("StartPprof alreadly in working...")
		return false
	}

	var err error

	// cpu profile
	for {
		_fcpu, err = os.Create(filename + ".cpu")
		if err != nil {
			fmt.Println("create CPU profile failed, err:", err)
			break
		}
		if err := pprof.StartCPUProfile(_fcpu); err != nil {
			fmt.Println("start CPU profile failed, err:", err)
			break
		}
		break
	}

	// memory profile
	for {
		f, err := os.Create(filename + ".memory")
		if err != nil {
			fmt.Println("create memory profile failed, err:", err)
		}
		runtime.GC()
		if err := pprof.Lookup("heap").WriteTo(f, 0); err != nil {
			// WriteHeapProfile(f) 等价于 Lookup("heap").WriteTo(w, 0)
			fmt.Println("write memory profile failed, err:", err)
		}
		f.Close()
		break
	}

	return _fcpu != nil
}

func ClosePprof() bool {
	if _fcpu != nil {
		pprof.StopCPUProfile()
		_fcpu.Close()
		_fcpu = nil
		return true
	}
	return false
}
