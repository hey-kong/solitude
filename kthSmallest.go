package main

import "container/heap"
import "strconv"

// Leetcode 230. (medium)
func kthSmallest(root *TreeNode, k int) int {
	_, res := recursiveKthSmallest(root, 0, k, -1)
	return res
}

func recursiveKthSmallest(root *TreeNode, i, k, res int) (int, int) {
	if root == nil || i >= k {
		return i, res
	}
	i, res = recursiveKthSmallest(root.Left, i, k, res)
	i++
	if i == k {
		return i, root.Val
	}
	i, res = recursiveKthSmallest(root.Right, i, k, res)
	return i, res
}

// Leetcode 5403. (hard)
func kthSmallest2(mat [][]int, k int) int {
	var h KthSmallestHeap
	heap.Init(&h)
	sum := 0
	for i := range mat {
		sum += mat[i][0]
	}
	idx := make([]int, len(mat))
	heap.Push(&h, KthSmallestPair{sum, idx})

	memo := make(map[string]bool)
	for i := 0; i < k; i++ {
		p := heap.Pop(&h).(KthSmallestPair)
		sum, idx = p.sum, p.idx
		for j := range idx {
			if idx[j] == len(mat[j])-1 {
				continue
			}
			tmpIdx := make([]int, len(idx))
			copy(tmpIdx, idx)
			tmpIdx[j]++
			str := toString(tmpIdx)
			if _, ok := memo[str]; !ok {
				tmpSum := sum - mat[j][idx[j]] + mat[j][idx[j]+1]
				heap.Push(&h, KthSmallestPair{tmpSum, tmpIdx})
				memo[str] = true
			}
		}
	}
	return sum
}

type KthSmallestPair struct {
	sum int
	idx []int
}

type KthSmallestHeap []KthSmallestPair

func (h KthSmallestHeap) Len() int {
	return len(h)
}

func (h KthSmallestHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h KthSmallestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KthSmallestHeap) Push(x interface{}) {
	*h = append(*h, x.(KthSmallestPair))
}

func (h *KthSmallestHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return tmp
}

func toString(arr []int) string {
	res := "["
	for i := 0; i < len(arr); i++ {
		res += strconv.Itoa(arr[i])
		if i != len(arr)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}
