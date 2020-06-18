package main

// Leetcode 386. (medium)
func lexicalOrder(n int) []int {
	res := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		res[i] = cur
		if cur*10 <= n {
			cur *= 10
		} else {
			if cur == n {
				cur /= 10
			}
			cur++
			for cur%10 == 0 {
				cur /= 10
			}
		}
	}
	return res
}
