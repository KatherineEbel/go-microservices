package utils

import (
	"sort"
)

func BubbleSort(slice []int) []int {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				keepRunning = true
			}
		}
	}
	return slice
}

func Sort(s []int) {
	if len(s) < 1000 {
		BubbleSort(s)
	}
	sort.Ints(s)
}
