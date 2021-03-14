package main

// Leetcode 5704. (hard)
func maximumScore(nums []int, k int) (res int) {
	m := nums[k]
	i, j := k, k
	for {
		for i > 0 && nums[i-1] >= m {
			i--
		}
		for j < len(nums)-1 && m <= nums[j+1] {
			j++
		}

		res = max(res, m*(j-i+1))
		if i > 0 && j < len(nums)-1 {
			m = max(nums[i-1], nums[j+1])
		} else if i > 0 {
			m = nums[i-1]
		} else if j < len(nums)-1 {
			m = nums[j+1]
		} else {
			break
		}
	}
	return res
}
