package main

import "container/heap"

// Leetcode 347. (medium)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	cntHeap := &CntHeap{}
	heap.Init(cntHeap)
	for num, cnt := range m {
		heap.Push(cntHeap, [2]int{num, cnt})
		if cntHeap.Len() > k {
			heap.Pop(cntHeap)
		}
	}
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(cntHeap).([2]int)[0]
	}
	return res
}

type CntHeap [][2]int

func (h CntHeap) Len() int {
	return len(h)
}

func (h CntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h CntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *CntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
