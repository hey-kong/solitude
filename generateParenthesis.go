package main

// Leetcode 22. (medium)
func generateParenthesis(n int) []string {
	return recursiveGenParenthesis(0, 0, n, "", []string{})
}

func recursiveGenParenthesis(left, right, n int, cur string, s []string) []string {
	if len(cur) == n * 2 {
		s = append(s, cur)
		return s
	}

	if left < n {
		s = recursiveGenParenthesis(left+1, right, n, cur+"(", s)
	}
	if right < left {
		s = recursiveGenParenthesis(left, right+1, n, cur+")", s)
	}
	return s
}
