package main

// Leetcode 994. (easy)
func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	mr, mc := len(grid), len(grid[0])
	memo := make([][]int, mr)
	for i := range memo {
		memo[i] = make([]int, mc)
	}
	cnt := 0
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	day := 0
	for len(queue) != 0 {
		locate := queue[0]
		queue = queue[1:]
		tmp := memo[locate[0]][locate[1]]
		if locate[0]-1 >= 0 && grid[locate[0]-1][locate[1]] == 1 {
			grid[locate[0]-1][locate[1]] = 2
			queue = append(queue, []int{locate[0] - 1, locate[1]})
			memo[locate[0]-1][locate[1]] = tmp + 1
			day = max(day, tmp+1)
			cnt--
		}
		if locate[1]-1 >= 0 && grid[locate[0]][locate[1]-1] == 1 {
			grid[locate[0]][locate[1]-1] = 2
			queue = append(queue, []int{locate[0], locate[1] - 1})
			memo[locate[0]][locate[1]-1] = tmp + 1
			day = max(day, tmp+1)
			cnt--
		}
		if locate[0]+1 < mr && grid[locate[0]+1][locate[1]] == 1 {
			grid[locate[0]+1][locate[1]] = 2
			queue = append(queue, []int{locate[0] + 1, locate[1]})
			memo[locate[0]+1][locate[1]] = tmp + 1
			day = max(day, tmp+1)
			cnt--
		}
		if locate[1]+1 < mc && grid[locate[0]][locate[1]+1] == 1 {
			grid[locate[0]][locate[1]+1] = 2
			queue = append(queue, []int{locate[0], locate[1] + 1})
			memo[locate[0]][locate[1]+1] = tmp + 1
			day = max(day, tmp+1)
			cnt--
		}
	}
	if cnt == 0 {
		return day
	}
	return -1
}
