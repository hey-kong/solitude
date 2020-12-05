package main

import "sort"

// Leetcode 621. (medium)
func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for _, b := range tasks {
		arr[b-'A']++
	}

	sort.Ints(arr)
	maxCnt := 0
	for _, num := range arr {
		if num == arr[25] {
			maxCnt++
		}
	}
	return max((arr[25]-1)*(n+1)+maxCnt, len(tasks))
}
