package main

// Leetcode m17.16. (easy)
func massage(nums []int) int {
	pre, cur := 0, 0
	for i := range nums {
		tmp := cur
		cur = max(cur, pre+nums[i])
		pre = tmp
	}
	return cur
}
