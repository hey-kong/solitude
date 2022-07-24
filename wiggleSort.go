package main

import "sort"

// Leetcode 324. (medium)
func wiggleSort(nums []int) {
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	n := len(arr)
	i, j, k := (n+1)/2-1, n-1, 0
	for i >= 0 {
		nums[k] = arr[i]
		k++
		i--
		if j >= (n+1)/2 {
			nums[k] = arr[j]
			k++
			j--
		}
	}
	return
}
