package main

// Leetcode 5892. (medium)
func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)
	for i := range stones {
		cnt[stones[i]%3]++
	}
	if cnt[0]%2 == 0 {
		return cnt[1] != 0 && cnt[2] != 0
	}
	return cnt[1] >= cnt[2]+3 || cnt[2] >= cnt[1]+3
}
