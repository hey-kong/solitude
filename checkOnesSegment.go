package main

// Leetcode 5697. (easy)
func checkOnesSegment(s string) bool {
	zero, one := 0, 0
	for i := range s {
		if s[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	prefix := 0
	for i := range s {
		if s[i] == '0' {
			break
		} else {
			prefix++
		}
	}
	return one == prefix
}
