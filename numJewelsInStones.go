package main

// Leetcode 771. (easy)
func numJewelsInStones(J string, S string) int {
	m := make(map[byte]bool)
	sum := 0
	for i := range J {
		m[J[i]] = true
	}
	for i := range S {
		if _, ok := m[S[i]]; ok {
			sum++
		}
	}
	return sum
}
