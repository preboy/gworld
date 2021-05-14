package poker

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestA45(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		fmt.Println("result:", random_angle_45(3, 7))
	}
}
