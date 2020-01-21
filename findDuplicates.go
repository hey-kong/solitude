package main

// Leetcode 442. (medium)
func findDuplicates(nums []int) []int {
	res := []int{}
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if j < 0 {
			i++
			continue
		}
		if nums[j] > 0 {
			nums[i] = nums[j]
			nums[j] = -1
		} else {
			nums[i] = 0
			nums[j]--
			if nums[j] == -2 {
                res = append(res, j+1)
			}
		}
	}
	return res
}
