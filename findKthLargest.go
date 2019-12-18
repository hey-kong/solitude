package main

import "math/rand"

// Leetcode 215. (medium)
func findKthLargest(nums []int, k int) int {
	l := len(nums)
	mark := 1

	// 消除对输入的依赖
	r := rand.Intn(l)
	pivot := nums[r]
	nums[r], nums[0] = nums[0], nums[r]
	for i := 1; i < l; i++ {
		if nums[i] < pivot {
			nums[mark], nums[i] = nums[i], nums[mark]
			mark++
		}
	}

	rightCount := l - mark

	// 右侧大于pivot的数量等于 k-1, 那pivot就是第k大的
	if rightCount == k - 1 {
		return pivot
	}

	// 在右侧继续查找
	if rightCount >= k {
		return findKthLargest(nums[mark:l], k)
	}

	// 在左侧继续查找
	return findKthLargest(nums[1:mark], k - rightCount - 1)
}
