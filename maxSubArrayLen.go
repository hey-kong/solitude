package main

// Leetcode 325. (medium)
func maxSubArrayLen(nums []int, k int) int {
	m := make(map[int]int, k)
	m[0] = -1
	maxLength := 0
	sum := 0
	for i, num := range nums {
		sum += num
		if j, ok := m[sum-k]; ok {
			maxLength = max(maxLength, i-j)
		}
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}
	return maxLength
}
