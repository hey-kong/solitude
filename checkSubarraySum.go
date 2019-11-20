package main

// Leetcode 523. (medium)
func checkSubarraySum1(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	for i := 0; i < len(sum) - 1; i++ {
		for j := i + 1; j < len(sum); j++ {
			res := sum[j] - sum[i] + nums[i]
			if k == 0 && res == 0 {
				return true
			}
			if k != 0 && res % k == 0 {
				return true
			}
		}
	}
	return false
}

func checkSubarraySum2(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	m := map[int]int{}
	m[0] = -1
	sum := 0
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}

		if j, ok := m[sum]; ok {
			if i - j != 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
