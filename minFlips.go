package main

// Leetcode 5778. (medium)
func minFlips(s string) int {
	target := "01"
	diff := 0
	for i := range s {
		if s[i] != target[i%2] {
			diff++
		}
	}

	res := min(diff, len(s)-diff)
	s += s
	for i := len(s) / 2; i < len(s); i++ {
		if s[i-len(s)/2] != target[(i-len(s)/2)%2] {
			diff--
		}
		if s[i] != target[i%2] {
			diff++
		}
		res = min(res, min(diff, len(s)/2-diff))
	}
	return res
}
