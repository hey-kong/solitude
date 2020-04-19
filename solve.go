package main

// Leetcode 130. (medium)
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, -1, 0, 1}
	r, c := len(board), len(board[0])
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}
	queue := [][2]int{}
	for i := 0; i < c; i++ {
		if board[0][i] == 'O' {
			visited[0][i] = true
			queue = append(queue, [2]int{0, i})
		}
	}
	for i := 0; i < r; i++ {
		if board[i][0] == 'O' {
			visited[i][0] = true
			queue = append(queue, [2]int{i, 0})
		}
	}
	for i := 0; i < r; i++ {
		if board[i][c-1] == 'O' {
			visited[i][c-1] = true
			queue = append(queue, [2]int{i, c - 1})
		}
	}
	for i := 0; i < c; i++ {
		if board[r-1][i] == 'O' {
			visited[r-1][i] = true
			queue = append(queue, [2]int{r - 1, i})
		}
	}

	for len(queue) > 0 {
		curX, curY := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x, y := curX+dx[i], curY+dy[i]
			if x >= 0 && x < r && y >= 0 && y < c && !visited[x][y] && board[x][y] == 'O' {
				visited[x][y] = true
				queue = append(queue, [2]int{x, y})
			}
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
