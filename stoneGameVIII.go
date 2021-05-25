package main

// Leetcode 1872. (hard)
func stoneGameVIII(stones []int) int {
	preSum := make([]int, len(stones))
	preSum[0] = stones[0]
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + stones[i]
	}

	res := preSum[len(preSum)-1]
	for i := len(preSum) - 2; i >= 1; i-- {
		res = max(res, preSum[i]-res)
	}
	return res
}
