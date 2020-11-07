package main

// Leetcode 316. (medium)
func removeDuplicateLetters(s string) string {
	ms := make(map[byte]int)
	mr := make(map[byte]bool)
	for i := range s {
		ms[s[i]] = i
	}
	res := ""
	for i := range s {
		if !mr[s[i]] {
			for res != "" && res[len(res)-1] > s[i] && ms[res[len(res)-1]] > i {
				mr[res[len(res)-1]] = false
				res = res[:len(res)-1]
			}
			mr[s[i]] = true
			res += string(s[i])
		}
	}
	return res
}
