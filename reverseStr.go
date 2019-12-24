package main

// Leetcode 541. (easy)
func reverseStr(s string, k int) string {
	b := []byte(s)
	for start := 0; start < len(s); start += 2*k {
		end := start + k
		if end > len(s) {
			end = len(s)
		}
		for j := start; j < (end + start) / 2; j++ {
			b[j], b[start+end-1-j] = b[start+end-1-j], b[j]
		}
	}
	return string(b)
}
