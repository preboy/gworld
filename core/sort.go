package core

import "sort"

func SortInt32s(arr []int32) {
	if len(arr) < 2 {
		return
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
