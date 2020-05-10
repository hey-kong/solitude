package main

// Leetcode 5407. (hard)
func ways(pizza []string, k int) int {
	mr, mc := len(pizza), len(pizza[0])
	// 前缀和数组nums
	nums := make([][]int, mr)
	for i := range nums {
		nums[i] = make([]int, mc)
	}
	if pizza[0][0] == 'A' {
		nums[0][0] = 1
	}
	for i := 1; i < mr; i++ {
		nums[i][0] = nums[i-1][0]
		if pizza[i][0] == 'A' {
			nums[i][0] += 1
		}
	}
	for i := 1; i < mc; i++ {
		nums[0][i] = nums[0][i-1]
		if pizza[0][i] == 'A' {
			nums[0][i] += 1
		}
	}
	for i := 1; i < mr; i++ {
		for j := 1; j < mc; j++ {
			nums[i][j] = nums[i-1][j] + nums[i][j-1] - nums[i-1][j-1]
			if pizza[i][j] == 'A' {
				nums[i][j] += 1
			}
		}
	}
	// 初始化dp
	dp := make([][][]int, mr)
	for i := range nums {
		dp[i] = make([][]int, mc)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	dp[0][0][1] = 1
	for x := 2; x <= k; x++ {
		for i := 0; i < mr; i++ {
			for j := 0; j < mc; j++ {
				if dp[i][j][x-1] == 0 {
					continue
				}
				// 水平
				for z := i + 1; z < mr; z++ {
					if waysHasA(nums, i, j, z-1, mc-1) && waysHasA(nums, z, j, mr-1, mc-1) {
						dp[z][j][x] += dp[i][j][x-1]
						dp[z][j][x] %= 1e9 + 7
					}
				}
				// 垂直
				for z := j + 1; z < mc; z++ {
					if waysHasA(nums, i, j, mr-1, z-1) && waysHasA(nums, i, z, mr-1, mc-1) {
						dp[i][z][x] += dp[i][j][x-1]
						dp[i][z][x] %= 1e9 + 7
					}
				}
			}
		}
	}
	res := 0
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			res += dp[i][j][k]
			res %= 1e9 + 7
		}
	}
	return res
}

func waysHasA(nums [][]int, sr, sc, er, ec int) bool {
	cnt1, cnt2, cnt3 := 0, 0, 0
	if sr-1 >= 0 && sc-1 >= 0 {
		cnt1 = nums[sr-1][sc-1]
	}
	if sr-1 >= 0 {
		cnt2 = nums[sr-1][ec]
	}
	if sc-1 >= 0 {
		cnt3 = nums[er][sc-1]
	}
	return nums[er][ec]-cnt2-cnt3+cnt1 > 0
}
