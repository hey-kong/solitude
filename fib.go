package main

// Leetcode 509. (easy)
func fib(n int) int {
	if n < 2 {
		return n
	}
	res := matrixPow([2][2]int{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func matrixPow(a [2][2]int, n int) [2][2]int {
	ret := [2][2]int{{1, 0}, {0, 1}}
	for n > 0 {
		if n&1 == 1 {
			ret = matrixMultiply(ret, a)
		}
		a = matrixMultiply(a, a)
		n >>= 1
	}
	return ret
}

func matrixMultiply(a, b [2][2]int) (c [2][2]int) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}
