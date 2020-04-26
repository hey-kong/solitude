package main

// Leetcode 5392. (easy)
func maxScore(s string) int {
	res := 0
	if s[0] == '0' {
		res++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			res++
		}
	}

	tmp := res
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '1' {
			tmp--
		}
		if s[i] == '0' {
			tmp++
			res = max(res, tmp)
		}
	}
	return res
}
