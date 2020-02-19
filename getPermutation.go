package main

// Leetcode 60. (medium)
func getPermutation(n int, k int) string {
	ori := "123456789"
	ori = ori[:n]
	res := ""
	k--
	for n > 0 {
		tmp := factorial(n - 1)
		i := k / tmp
		res += ori[i : i+1]
		ori = ori[:i] + ori[i+1:]
		k %= tmp
		n--
	}
	return res
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
