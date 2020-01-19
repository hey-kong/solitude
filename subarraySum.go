package main

// Leetcode 560. (medium)
func subarraySum(nums []int, k int) int {
	m := make(map[int]int, k)
	m[0] = 1
	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if i, ok := m[sum-k]; ok {
			count += i
		}
		m[sum]++
	}
	return count
}
