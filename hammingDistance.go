package main

// Leetcode 461. (easy)
func hammingDistance(x int, y int) int {
	pos := x ^ y
	cnt := 0
	for pos != 0 {
        pos &= (pos-1)
		cnt++
	}
	return cnt
}
