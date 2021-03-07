package main

// Leetcode 5698. (medium)
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	c := goal - sum
	if c < 0 {
		c = -c
	}
	if c%limit == 0 {
		return c / limit
	}
	return c/limit + 1
}
