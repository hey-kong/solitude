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
	for {
		// 先判断，再移动指针
		// 因为重复数字可能是nums[0]
		if slow == fast {
			break
		}
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
