package lobby

import (
	"fmt"
	"testing"
)

func TestAnalyse(t *testing.T) {
	cards := []Card{15002, 15003, 15000, 15001, 12001}

	a := NewAnalyse(cards)

	for _, v := range a.Points() {
		fmt.Println(v.Point, v.Count)
	}

}
