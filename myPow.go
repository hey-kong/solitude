package main

// Leetcode 50. (medium)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}
	return res
}
