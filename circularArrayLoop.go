package main

// Leetcode 457. (medium)
func circularArrayLoop(nums []int) bool {
	for i, num := range nums {
		if num == 0 {
			continue
		}
		slow, fast := i, nextCircularArrayLoop(i, nums)
		for num*nums[fast] > 0 && num*nums[nextCircularArrayLoop(fast, nums)] > 0 {
			if slow == fast {
				if slow == nextCircularArrayLoop(slow, nums) {
					break
				}
				return true
			}
			slow = nextCircularArrayLoop(slow, nums)
			fast = nextCircularArrayLoop(nextCircularArrayLoop(fast, nums), nums)
		}
		slow = i
		for num*nums[slow] > 0 {
			next := nextCircularArrayLoop(slow, nums)
			nums[slow] = 0
			slow = next
		}
	}
	return false
}

func nextCircularArrayLoop(i int, nums []int) int {
	return ((i+nums[i])%len(nums) + len(nums)) % len(nums)
}
