package main

// Leetcode 264. (medium)
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		num := min(min(nums[i2]*2, nums[i3]*3), nums[i5]*5)
		nums[i] = num
		if nums[i2]*2 == num {
			i2++
		}
		if nums[i3]*3 == num {
			i3++
		}
		if nums[i5]*5 == num {
			i5++
		}
	}
	return nums[n-1]
}
