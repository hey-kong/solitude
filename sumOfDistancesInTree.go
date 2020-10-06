package main

// Leetcode 834. (hard)
var resOfDistancesInTree []int
var cntOfDistancesInTree []int
var mapOfDistancesInTree []map[int]bool

func sumOfDistancesInTree(N int, edges [][]int) []int {
	resOfDistancesInTree = make([]int, N)
	cntOfDistancesInTree = make([]int, N)
	mapOfDistancesInTree = make([]map[int]bool, N)
	for i := range mapOfDistancesInTree {
		mapOfDistancesInTree[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		mapOfDistancesInTree[edge[0]][edge[1]] = true
		mapOfDistancesInTree[edge[1]][edge[0]] = true
	}
	recursiveSumOfDistancesInTree(0, -1)
	recursiveSumOfDistancesInTree2(0, -1)
	return resOfDistancesInTree
}

func recursiveSumOfDistancesInTree(root, father int) {
	for i := range mapOfDistancesInTree[root] {
		if i == father {
			continue
		}
		recursiveSumOfDistancesInTree(i, root)
		cntOfDistancesInTree[root] += cntOfDistancesInTree[i]
		resOfDistancesInTree[root] += resOfDistancesInTree[i] + cntOfDistancesInTree[i]
	}
	cntOfDistancesInTree[root]++
}

func recursiveSumOfDistancesInTree2(root, father int) {
	for i := range mapOfDistancesInTree[root] {
		if i == father {
			continue
		}
		resOfDistancesInTree[i] = resOfDistancesInTree[root] - cntOfDistancesInTree[i] + (len(cntOfDistancesInTree) - cntOfDistancesInTree[i])
		recursiveSumOfDistancesInTree2(i, root)
	}
}
