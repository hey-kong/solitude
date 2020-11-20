package main

// Leetcode 91. (medium)
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' && s[i-1] != '1' && s[i-1] != '2' {
			return 0
		} else if s[i] == '0' {
			pre, cur = cur, pre
		} else if (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1') {
			pre, cur = cur, cur+pre
		} else {
			pre, cur = cur, cur
		}
	}
	return cur
}
