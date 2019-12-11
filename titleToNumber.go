package main

// Leetcode 171. (easy)
func titleToNumber(s string) int {
	res := 0
	pos := 1
	l := len(s)
	for l > 0 {
		l--
		res += (int(s[l] - 'A') + 1) * pos
		pos *= 26
	}
	return res
}
