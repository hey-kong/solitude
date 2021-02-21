package main

import "container/heap"

// Leetcode 767. (medium)
func reorganizeString(S string) string {
	m := make(map[byte]int)
	for i := range S {
		m[S[i]]++
	}

	h := ByteCntHeap{}
	heap.Init(&h)
	for k, v := range m {
		heap.Push(&h, ByteCnt{k, v})
	}
	if h[0].cnt > (len(S)+1)/2 {
		return ""
	}

	res := []byte(S)
	i := 0
	for h.Len() != 0 {
		b, cnt := h[0].b, h[0].cnt
		heap.Pop(&h)
		for j := 0; j < cnt; j++ {
			res[i] = b
			i += 2
			if i >= len(S) {
				i = 1
			}
		}
	}
	return string(res)
}

type ByteCnt struct {
	b   byte
	cnt int
}

type ByteCntHeap []ByteCnt

func (h ByteCntHeap) Len() int {
	return len(h)
}

func (h ByteCntHeap) Less(i, j int) bool {
	return h[i].cnt > h[j].cnt
}

func (h ByteCntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ByteCntHeap) Push(x interface{}) {
	*h = append(*h, x.(ByteCnt))
}

func (h *ByteCntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
