package main

// Leetcode 1744. (medium)
func canEat(candiesCount []int, queries [][]int) []bool {
	preSum := make([]int, len(candiesCount))
	preSum[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++ {
		preSum[i] = preSum[i-1] + candiesCount[i]
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		favoriteType, favoriteDay, dailyCap := q[0], q[1], q[2]
		mx := (favoriteDay + 1) * dailyCap
		mn := favoriteDay + 1
		if favoriteType == 0 {
			if mn <= preSum[0] {
				ans[i] = true
			}
		} else {
			if mx > preSum[favoriteType-1] && mn <= preSum[favoriteType] {
				ans[i] = true
			}
		}
	}
	return ans
}
