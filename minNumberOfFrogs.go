package main

// Leetcode 5390. (medium)
func minNumberOfFrogs(croakOfFrogs string) int {
	c, r, o, a, k := 0, 0, 0, 0, 0
	res := 0
	for i := range croakOfFrogs {
		if croakOfFrogs[i] == 'c' {
			c++
		} else if croakOfFrogs[i] == 'r' {
			r++
		} else if croakOfFrogs[i] == 'o' {
			o++
		} else if croakOfFrogs[i] == 'a' {
			a++
		} else if croakOfFrogs[i] == 'k' {
			k++
		}
		if c < r || r < o || o < a || a < k {
			return -1
		}
		res = max(res, c-k)
	}
	if c != r || c != o || c != a || c != k {
		return -1
	}
	return res
}
