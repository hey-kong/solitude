package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

// Leetcode 5703. (medium)
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := PassCntHeap{}
	for _, c := range classes {
		heap.Push(&h, [2]int{c[0], c[1]})
	}
	for i := 0; i < extraStudents; i++ {
		c := heap.Pop(&h).([2]int)
		c[0]++
		c[1]++
		heap.Push(&h, [2]int{c[0], c[1]})
	}
	total := 0.0
	for _, c := range h {
		total += divide5(c[0], c[1])
	}
	return total / float64(h.Len())
}

func divide5(a, b int) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(a)/float64(b)), 64) // 保留5位小数
	return res
}

type PassCntHeap [][2]int

func (h PassCntHeap) Len() int {
	return len(h)
}

func (h PassCntHeap) Less(i, j int) bool {
	return float64(h[i][1]-h[i][0])/float64((h[i][1]+1)*h[i][1]) > float64(h[j][1]-h[j][0])/float64((h[j][1]+1)*h[j][1])
}

func (h PassCntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PassCntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *PassCntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
