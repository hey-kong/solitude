package main

// Leetcode 327. (hard)
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	aux := make([]int, len(nums))
	return mergeSortOfCountRangeSum(preSum, aux, 0, len(nums)-1, lower, upper)
}

func mergeSortOfCountRangeSum(preSum, aux []int, left, right int, lower int, upper int) int {
	if left == right {
		if preSum[left] <= upper && preSum[left] >= lower {
			return 1
		}
		return 0
	}
	mid := left + (right-left)/2
	ans1 := mergeSortOfCountRangeSum(preSum, aux, left, mid, lower, upper)
	ans2 := mergeSortOfCountRangeSum(preSum, aux, mid+1, right, lower, upper)
	i, j1, j2 := left, mid+1, mid+1
	ans3 := 0
	for i < mid+1 {
		for j1 <= right && preSum[j1]-preSum[i] < lower {
			j1++
		}
		j2 = j1
		for j2 <= right && preSum[j2]-preSum[i] <= upper {
			j2++
			ans3++
		}
		i++
	}
	mergeArray(preSum, aux, left, mid, right)
	return ans1 + ans2 + ans3
}
