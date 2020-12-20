package main

// Leetcode 1248. (medium)
func numberOfSubarrays(nums []int, k int) int {
	sum, res := 0, 0
	prefixCnt := make([]int, len(nums)+1)
	prefixCnt[0] = 1
	for _, num := range nums {
		sum += num & 1
		prefixCnt[sum]++
		if sum >= k {
			res += prefixCnt[sum-k]
		}
	}
	return res
}
