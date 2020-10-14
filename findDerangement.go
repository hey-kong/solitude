package main

// Leetcode 634. (medium)
func findDerangement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	a, b := 0, 1
	res := 0
	for i := 3; i <= n; i++ {
		res = ((i - 1) * (a + b)) % 1000000007
		a = b
		b = res
	}
	return res
}
