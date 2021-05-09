package main

// Leetcode 5752. (medium)
func maxSumMinProduct(nums []int) int {
	nums = append([]int{0}, nums...)
	nums = append(nums, 0)
	prefix := make([]int, len(nums))
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	stack := []int{}
	right := make([]int, len(nums))
	for i := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] > nums[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	left := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] < nums[stack[len(stack)-1]] {
			left[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 1; i < len(nums)-1; i++ {
		res = max(res, nums[i]*(prefix[right[i]-1]-prefix[left[i]]))
	}
	return res % 1000000007
}
