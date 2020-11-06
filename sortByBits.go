package main

import "sort"

// Leetcode 1356. (easy)
func sortByBits(arr []int) []int {
	bitsOfNum := make([]int, 10001)
	for i := 1; i < 10001; i++ {
		bitsOfNum[i] = bitsOfNum[i>>1] + (i & 1)
	}
	sort.Slice(arr, func(i, j int) bool {
		if bitsOfNum[arr[i]] == bitsOfNum[arr[j]] {
			return arr[i] < arr[j]
		}
		return bitsOfNum[arr[i]] < bitsOfNum[arr[j]]
	})
	return arr
}
