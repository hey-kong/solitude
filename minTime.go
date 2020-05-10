package main

// Leetcode 5406. (medium)
func minTime(n int, edges [][]int, hasApple []bool) int {
	notLeaf := make([]bool, n)
	edgeTo := make([]int, n)
	for _, e := range edges {
		edgeTo[e[1]] = e[0]
		notLeaf[e[0]] = true
	}
	leafs := []int{}
	for i, flag := range notLeaf {
		if !flag {
			leafs = append(leafs, i)
		}
	}
	for _, leaf := range leafs {
		for i := leaf; i != 0; i = edgeTo[i] {
			if hasApple[i] {
				hasApple[edgeTo[i]] = true
			}
		}
	}
	res := 0
	for i := 1; i < n; i++ {
		if hasApple[i] {
			res++
		}
	}
	return res * 2
}
