package main

// Leetcode 207. (medium)
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, pair := range prerequisites {
		adj[pair[1]] = append(adj[pair[1]], pair[0])
	}
	flags := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		finish := false
		finish, flags = canFinishDFS(i, adj, flags)
		if !finish {
			return false
		}
	}
	return true
}

func canFinishDFS(i int, adj [][]int, flags []int) (bool, []int) {
	if flags[i] == 1 {
		return false, flags
	} else if flags[i] == -1 {
		return true, flags
	}

	flags[i] = 1
	for _, j := range adj[i] {
		finish := false
		finish, flags = canFinishDFS(j, adj, flags)
		if !finish {
			return false, flags
		}
	}
	flags[i] = -1
	return true, flags
}

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	pre := make([]int, numCourses)
	for _, pair := range prerequisites {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
		pre[pair[0]]++
	}
	queue := []int{}
	for i := range adj {
		if pre[i] == 0 {
			queue = append(queue, i)
			numCourses--
		}
	}
	for len(queue) != 0 {
		for i := range adj {
			if i == queue[0] {
				continue
			}
			for j := range adj[i] {
				if adj[i][j] == queue[0] {
					adj[i][j] = -1
					pre[i]--
					if pre[i] == 0 {
						queue = append(queue, i)
						numCourses--
					}
					break
				}
			}
		}
		queue = queue[1:]
	}
	return numCourses == 0
}
