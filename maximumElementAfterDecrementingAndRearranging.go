package main

import "sort"

// Leetcode 1846. (medium)
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	arr[0] = 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1]-1 {
			arr[i+1] = arr[i] + 1
		}
	}
	return arr[len(arr)-1]
}
