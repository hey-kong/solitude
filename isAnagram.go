package main

// Leetcode 242. (easy)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make([]int, 26)
	for _, c := range s {
		table[c-'a']++
	}
	for _, c := range t {
		table[c-'a']--
		if table[c-'a'] < 0 {
			return false
		}
	}
	return true
}
