package main

// Leetcode 1. (easy)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			res[0] = i
			res[1] = j
			break
		}
		m[num] = i
	}
	return res
}
