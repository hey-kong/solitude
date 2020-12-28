package main

// Leetcode 187. (medium)
func findRepeatedDnaSequences(s string) (res []string) {
	L := 10
	m, mRes := make(map[string]bool), make(map[string]bool)
	for i := 0; i <= len(s)-L; i++ {
		if m[s[i:i+L]] {
			mRes[s[i:i+L]] = true
		} else {
			m[s[i:i+L]] = true
		}
	}
	for str := range mRes {
		res = append(res, str)
	}
	return res
}
