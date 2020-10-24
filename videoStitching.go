package main

// Leetcode 1024. (medium)
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := 1; i <= T; i++ {
		dp[i] = -1
	}
	for i := 1; i <= T; i++ {
		for _, c := range clips {
			if c[0] <= i && i <= c[1] {
				if dp[c[0]] != -1 {
					if dp[i] == -1 {
						dp[i] = dp[c[0]] + 1
					} else {
						dp[i] = min(dp[i], dp[c[0]]+1)
					}
				}
			}
		}
	}
	return dp[T]
}

func videoStitching2(clips [][]int, T int) int {
	// switch to jump game
	nums := make([]int, T)
	for _, c := range clips {
		if c[0] < T {
			nums[c[0]] = max(nums[c[0]], c[1])
		}
	}

	step := 0
	end := 0
	lastpos := 0
	for i, num := range nums {
		lastpos = max(lastpos, num)
		if i == end {
			if end == lastpos {
				return -1
			}
			step++
			end = lastpos
		}
	}
	return step
}
