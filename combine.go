package main

// Leetcode 77. (medium)
func combine(n int, k int) [][]int {
	return dfsCombine(1, n, k, []int{}, [][]int{})
}

func dfsCombine(i, n, k int, cur []int, res [][]int) [][]int {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		res = append(res, tmp)
		return res
	}

	for j := i; j <= n; j++ {
		cur = append(cur, j)
		res = dfsCombine(j+1, n, k, cur, res)
		cur = cur[:len(cur)-1]
	}
	return res
}
