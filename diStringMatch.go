package main

// Leetcode 942. (easy)
func diStringMatch(S string) []int {
	nums := []int{}
	left, right := 0, len(S)
	for _, c := range S {
		if c == 'I' {
			nums = append(nums, left)
			left++
		}
		if c == 'D' {
			nums = append(nums, right)
			right--
		}
	}
	nums = append(nums, left)
	return nums
}
