package main

// Leetcode 254. (medium)
func getFactors(n int) [][]int {
	return dfsGetFactors(n, 2)
}

func dfsGetFactors(n, start int) [][]int {
	res := [][]int{}
	for i := start; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, []int{n / i, i})

			for _, cur := range dfsGetFactors(n/i, i) {
				cur = append(cur, i)
				res = append(res, cur)
			}
		}
	}
	return res
}
