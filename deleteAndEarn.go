package main

// Leetcode 740. (medium)
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}

	arr := make([]int, maxVal+1)
	for _, num := range nums {
		arr[num]++
	}

	pre, cur := 0, 0
	for i, num := range arr {
		pre, cur = cur, max(pre+i*num, cur)
	}
	return cur
}
