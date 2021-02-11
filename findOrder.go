package main

// Leetcode 210. (medium)
func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
	in := make([]int, numCourses)
	queue := []int{}

	for _, p := range prerequisites {
		in[p[0]]++
	}
	for i := range in {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		tmp := queue[0]
		queue = queue[1:]
		res = append(res, tmp)
		for _, p := range prerequisites {
			if p[1] == tmp {
				in[p[0]]--
				if in[p[0]] == 0 {
					queue = append(queue, p[0])
				}
			}
		}
	}

	if len(res) != numCourses {
		return nil
	}
	return res
}
