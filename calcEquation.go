package main

// Leetcode 399. (medium)
func calcEquation(equations [][]string, values []float64, queries [][]string) (res []float64) {
	id := make(map[string]int)
	for i := range equations {
		if _, ok := id[equations[i][0]]; !ok {
			id[equations[i][0]] = len(id)
		}
		if _, ok := id[equations[i][1]]; !ok {
			id[equations[i][1]] = len(id)
		}
	}

	graph := make([][]float64, len(id))
	for i := range graph {
		graph[i] = make([]float64, len(id))
	}
	for i := range values {
		graph[id[equations[i][0]]][id[equations[i][1]]] = values[i]
		graph[id[equations[i][1]]][id[equations[i][0]]] = 1.0 / values[i]
	}

	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] != 0 && graph[k][j] != 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	res = make([]float64, len(queries))
	for i, query := range queries {
		start, hasS := id[query[0]]
		end, hasE := id[query[1]]
		if !hasS || !hasE || graph[start][end] == 0 {
			res[i] = -1
		} else {
			res[i] = graph[start][end]
		}
	}
	return
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) (res []float64) {
	id := make(map[string]int)
	for i := range equations {
		if _, ok := id[equations[i][0]]; !ok {
			id[equations[i][0]] = len(id)
		}
		if _, ok := id[equations[i][1]]; !ok {
			id[equations[i][1]] = len(id)
		}
	}

	root, w := make([]int, len(id)), make([]float64, len(id))
	for i := 0; i < len(id); i++ {
		root[i] = i
		w[i] = 1.0
	}
	for i, v := range values {
		root, w = unionOf399(id[equations[i][0]], id[equations[i][1]], v, root, w)
	}

	res = make([]float64, len(queries))
	for i, query := range queries {
		x, hasX := id[query[0]]
		y, hasY := id[query[1]]
		if !hasX || !hasY {
			res[i] = -1.0
			continue
		}
		r1, w1 := findOf399(x, root, w)
		r2, w2 := findOf399(y, root, w)
		if r1 != r2 {
			res[i] = -1.0
		} else {
			res[i] = w1 / w2
		}
	}
	return
}

func findOf399(x int, root []int, w []float64) (int, float64) {
	v := 1.0
	for x != root[x] {
		v *= w[x]
		x = root[x]
	}
	return x, v
}

func unionOf399(x, y int, v float64, root []int, w []float64) ([]int, []float64) {
	r1, w1 := findOf399(x, root, w)
	r2, w2 := findOf399(y, root, w)
	if r1 != r2 {
		root[r2] = r1
		w[r2] = w1 / (w2 * v)
	}
	return root, w
}
