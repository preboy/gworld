package rand

import (
	"math/rand"
	"time"
)

func SetSeed() {
	rand.Seed(time.Now().Unix())
}

// returns true if x <= prob where x is a random number in [0, n)
func RandomHitn(prob, n int) bool {
	return rand.Intn(n) < prob
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

func Uint32(n int32) uint32 {
	return uint32(rand.Int31n(n))
}

func Int32(n int32) int32 {
	return rand.Int31n(n)
}
