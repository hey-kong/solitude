package main

// Leetcode 440. (hard)
func findKthNumber(n int, k int) int {
	i := 1
	prefix := 1
	for i < k {
		cnt := getCnt(prefix, n)
		if i+cnt > k {
			i++
			prefix *= 10
		} else if i+cnt <= k {
			i += cnt
			prefix++
		}
	}
	return prefix
}

func getCnt(prefix, n int) int {
	cur, next := prefix, prefix+1
	cnt := 0
	for cur <= n {
		cnt += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}
	return cnt
}
