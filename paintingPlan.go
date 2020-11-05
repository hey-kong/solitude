package main

// Leetcode LCP.22. (easy)
func paintingPlan(n int, k int) int {
	if k == n*n {
		return 1
	}
	res := 0
	for i := 0; i < n; i++ {
		j := (k - n*i) / (n - i)
		if j*(n-i) == (k-n*i) && j >= 0 && j < n {
			res += combination(n, i) * combination(n, j)
		}
	}
	return res
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func combination(x, y int) int {
	return fact(x) / (fact(y) * fact(x-y))
}
