package math

import (
	"math/rand"
)

// returns true if x <= prob where x is a random number in [0, n)
func RandomHitn(prob, n int) bool {
	r := rand.Intn(n)
	if r < prob {
		return true
	}
	return false
}

// returns [a, b]
func RandomInt(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		a, b = b, a
	}

	return rand.Intn(b-a+1) + a
}
