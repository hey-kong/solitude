package main

// Leetcode 389. (easy)
func findTheDifference(s string, t string) byte {
	res := t[0]
	for i := range s {
		res ^= s[i]
		res ^= t[i+1]
	}
	return res
}
