package main

// Leetcode 300. (medium)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := make([]int, 1)
	res[0] = nums[0]
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > res[j] {
			res = append(res, nums[i])
			j++
			continue
		}
		left, right := 0, j
		for left <= right {
			mid := left + (right-left)/2
			if res[mid] == nums[i] {
				left = mid
				break
			} else if res[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		res[left] = nums[i]
	}
	return j + 1
}
