package main

// Leetcode 287. (medium)
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = nums[0]
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}