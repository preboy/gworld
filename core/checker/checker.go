package checker

type checker = func(uint32, ...interface{}) bool

var (
	_checkers = make(map[uint32][]checker, 0x100)
)

// ----------------------------------------------------------------------------

func Verify(id uint32, fn checker) {
	_checkers[id] = append(_checkers[id], fn)
}

func Any(id uint32, args ...interface{}) bool {
	if len(_checkers[id]) == 0 {
		return true
	}

	for _, fn := range _checkers[id] {
		if fn(id, args...) {
			return true
		}
	}

	return false
}

func All(id uint32, args ...interface{}) bool {
	if len(_checkers[id]) == 0 {
		return true
	}

	for _, fn := range _checkers[id] {
		if !fn(id, args...) {
			return false
		}
	}

	return true
}
