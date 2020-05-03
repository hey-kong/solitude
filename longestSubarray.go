package main

// Leetcode 5402. (medium)
func longestSubarray(nums []int, limit int) int {
	maxStack := []int{}
	minStack := []int{}
	left := 0
	res := 0
	for right, num := range nums {
		for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] < num {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, right)
		for len(minStack) > 0 && nums[minStack[len(minStack)-1]] > num {
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, right)

		for len(maxStack) > 0 && len(minStack) > 0 && nums[maxStack[0]]-nums[minStack[0]] > limit {
			left++
			if left > maxStack[0] {
				maxStack = maxStack[1:]
			}
			if left > minStack[0] {
				minStack = minStack[1:]
			}
		}
		res = max(res, right-left+1)
	}
	return res
}
