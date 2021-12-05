package main

// Leetcode 372. (medium)
func superPow(a int, b []int) int {
	return dfsSuperPow(a, b, len(b)-1)
}

func dfsSuperPow(a int, b []int, last int) int {
	if last < 0 {
		return 1
	}
	return subSuperPow(dfsSuperPow(a, b, last-1), 10) * subSuperPow(a, b[last]) % 1337
}

func subSuperPow(a int, b int) int {
	res := 1
	a %= 1337
	for b != 0 {
		if b&1 == 1 {
			res *= a
		}
		a = a * a % 1337
		b >>= 1
	}
	return res
}
