package main

// Leetcode 32. (hard)
func longestValidParentheses(s string) int {
	stack := []int{}
	dp := make([]int, len(s))
	maxLength := 0
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' && len(stack) > 0 {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if j == 0 {
				dp[i] = i - j + 1
			} else {
				dp[i] = i - j + 1 + dp[j-1]
			}
			maxLength = max(maxLength, dp[i])
		}
	}
	return maxLength
}
