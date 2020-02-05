package main

import "sort"

type Triples [][]int

// Leetcode 987. (medium)
func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	triples := recursiveVerticalTraversal(root, 0, 0, [][]int{})
	sort.Sort(Triples(triples))
	res := [][]int{}
	preX := -1 << 31
	tmp := []int{}
	for _, triple := range triples {
		if triple[0] != preX && preX != -1<<31 {
			arr := make([]int, len(tmp))
			copy(arr, tmp)
			res = append(res, arr)
			tmp = tmp[:0]
		}
		tmp = append(tmp, triple[2])
		preX = triple[0]
	}
	arr := make([]int, len(tmp))
	copy(arr, tmp)
	res = append(res, arr)
	return res
}

func recursiveVerticalTraversal(root *TreeNode, x, y int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	triple := make([]int, 3)
	triple[0], triple[1], triple[2] = x, y, root.Val
	res = append(res, triple)
	res = recursiveVerticalTraversal(root.Left, x-1, y-1, res)
	res = recursiveVerticalTraversal(root.Right, x+1, y-1, res)
	return res
}

func (arr Triples) Len() int {
	return len(arr)
}

func (arr Triples) Less(i, j int) bool {
	if arr[i][0] == arr[j][0] && arr[i][1] == arr[j][1] {
		return arr[i][2] < arr[j][2]
	}
	if arr[i][0] == arr[j][0] {
		return arr[i][1] > arr[j][1]
	}
	return arr[i][0] < arr[j][0]
}

func (arr Triples) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
