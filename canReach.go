package main

// Leetcode 1871. (medium)
func canReach(s string, minJump int, maxJump int) bool {
	dp := make([]bool, len(s))
	dp[0] = true
	cnt := 1
	for i := minJump; i < len(s); i++ {
		if s[i] == '0' && cnt > 0 {
			dp[i] = true
		}
		if i-maxJump >= 0 && dp[i-maxJump] {
			cnt--
		}
		if dp[i-minJump+1] {
			cnt++
		}
	}
	return dp[len(s)-1]
}
