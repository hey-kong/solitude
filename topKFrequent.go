package main

import "container/heap"

// Leetcode 347. (medium)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	h := &CntHeap{}
	heap.Init(h)
	for num, cnt := range m {
		heap.Push(h, [2]int{num, cnt})
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = (heap.Pop(h).([2]int))[0]
	}
	return res
}

type CntHeap [][2]int

func (h CntHeap) Len() int {
	return len(h)
}

func (h CntHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
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

// Leetcode 692. (medium)
func topKFrequent2(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	h := &WordCntHeap{}
	heap.Init(h)
	for word, cnt := range m {
		heap.Push(h, WordCnt{word, cnt})
	}
	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(WordCnt).word
	}
	return res
}

type WordCnt struct {
	word string
	cnt  int
}

type WordCntHeap []WordCnt

func (h WordCntHeap) Len() int {
	return len(h)
}

func (h WordCntHeap) Less(i, j int) bool {
	if h[i].cnt == h[j].cnt {
		return h[i].word < h[j].word
	}
	return h[i].cnt > h[j].cnt
}

func (h WordCntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordCntHeap) Push(x interface{}) {
	*h = append(*h, x.(WordCnt))
}

func (h *WordCntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
