package main

// Leetcode 647. (medium)
func countSubstrings(s string) int {
	res := 0
	isPalindrome := make([][]bool, len(s))
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, len(s))
	}
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			if s[i] == s[i+l-1] && (l <= 3 || isPalindrome[i+1][i+l-2]) {
				res++
				isPalindrome[i][i+l-1] = true
			}
		}
	}
	return res
}
