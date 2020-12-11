package main

// Leetcode 1198. (medium)
func smallestCommonElement(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	idx := make([]int, m)
	maxNum := mat[0][0]
	cnt := 0
	for {
		for i := 0; i < m; i++ {
			left, right := idx[i], n-1
			for left <= right {
				mid := left + (right-left)/2
				if mat[i][mid] == maxNum {
					left = mid
					cnt++
					break
				} else if mat[i][mid] < maxNum {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			if cnt == m {
				return maxNum
			}
			if left == n {
				return -1
			}
			idx[i] = left
			if mat[i][left] != maxNum {
				maxNum = mat[i][idx[i]]
				cnt = 1
			}
		}
	}
}
