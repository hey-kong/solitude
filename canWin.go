package main

// Leetcode 294. (medium)
func canWin(currentState string) bool {
	s := []byte(currentState)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			s[i] = '-'
			s[i+1] = '-'
			if !canWin(string(s)) {
				return true
			}
			s[i] = '+'
			s[i+1] = '+'
		}
	}
	return false
}
