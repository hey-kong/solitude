package main

// Leetcode 205. (easy)
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		if s2t[s[i]] == 0 && t2s[t[i]] == 0 {
			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		} else if s2t[s[i]] != t[i] || t2s[t[i]] != s[i] {
			return false
		}
	}
	return true
}
