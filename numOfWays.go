package main

// Leetcode 1411. (hard)
func numOfWays(n int) int {
	aba, abc := 6, 6
	offset := 1000000007
	for i := 1; i < n; i++ {
		aba, abc = aba*3+abc*2, aba*2+abc*2
		aba %= offset
		abc %= offset
	}
	return (aba + abc) % offset
}
