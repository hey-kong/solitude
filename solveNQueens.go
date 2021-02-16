package main

import "strings"

// Leetcode 51. (hard)
func solveNQueens(n int) [][]string {
	empty := strings.Repeat(".", n)
	cur := make([]string, n)
	for i := range cur {
		cur[i] = empty
	}
	res := dfsSolveNQueens(0, n, cur, [][]string{})
	return res
}

func dfsSolveNQueens(i, n int, cur []string, res [][]string) [][]string {
	if i == n {
		tmp := make([]string, n)
		copy(tmp, cur)
		res = append(res, tmp)
		return res
	}

	for j := range cur[i] {
		cur[i] = cur[i][:j] + "Q" + cur[i][j+1:]
		if isValidOfSolveNQueens(n, i, j, cur) {
			res = dfsSolveNQueens(i+1, n, cur, res)
		}
		cur[i] = cur[i][:j] + "." + cur[i][j+1:]
	}
	return res
}

func isValidOfSolveNQueens(n, row, col int, cur []string) bool {
	for i := 0; i < row; i++ {
		if cur[i][col] == 'Q' {
			return false
		}
	}
	i, j := row, col
	for i > 0 && j > 0 {
		i--
		j--
		if cur[i][j] == 'Q' {
			return false
		}
	}
	i, j = row, col
	for i > 0 && j < n-1 {
		i--
		j++
		if cur[i][j] == 'Q' {
			return false
		}
	}
	return true
}
