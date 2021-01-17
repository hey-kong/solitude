package main

// Leetcode 161. (medium)
func isOneEditDistance(s string, t string) bool {
	ns, nt := len(s), len(t)
	if ns < nt {
		return isOneEditDistance(t, s)
	}
	if ns-nt > 1 || s == t {
		return false
	}

	for i := range t {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i+1:] == t[i:]
			}
		}
	}
	return true
}
