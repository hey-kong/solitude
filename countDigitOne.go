package main

// Leetcode 233. (hard)
func countDigitOne(n int) (res int) {
	H := n / 10
	i := n % 10
	L := 0
	base := 1
	for H != 0 || i != 0 {
		res += H * base
		if i > 1 {
			res += base
		} else if i == 1 {
			res += L + 1
		}

		L += i * base
		i = H % 10
		H /= 10
		base *= 10
	}
	return
}
