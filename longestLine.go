package main

// Leetcode 562. (medium)
func longestLine(mat [][]int) (res int) {
	horizontal := make([]int, len(mat[0]))
	vertical := make([]int, len(mat[0]))
	diagonal := make([]int, len(mat[0]))
	antidiagonal := make([]int, len(mat[0]))
	for i := range mat {
		tmpVertical := make([]int, len(mat[0]))
		tmpDiagonal := make([]int, len(mat[0]))
		tmpAntidiagonal := make([]int, len(mat[0]))
		for j := range mat[i] {
			if mat[i][j] == 0 {
				horizontal[j], tmpVertical[j], tmpDiagonal[j], tmpAntidiagonal[j] = 0, 0, 0, 0
			} else {
				if j == 0 {
					horizontal[j] = 1
				} else {
					horizontal[j] = horizontal[j-1] + 1
				}

				if i == 0 {
					tmpVertical[j] = 1
				} else {
					tmpVertical[j] = vertical[j] + 1
				}

				if i == 0 || j == 0 {
					tmpDiagonal[j] = 1
				} else {
					tmpDiagonal[j] = diagonal[j-1] + 1
				}

				if i == 0 || j == len(mat[0])-1 {
					tmpAntidiagonal[j] = 1
				} else {
					tmpAntidiagonal[j] = antidiagonal[j+1] + 1
				}
			}
			res = max(res, horizontal[j])
			res = max(res, tmpVertical[j])
			res = max(res, tmpDiagonal[j])
			res = max(res, tmpAntidiagonal[j])
		}
		vertical, diagonal, antidiagonal = tmpVertical, tmpDiagonal, tmpAntidiagonal
	}
	return
}
