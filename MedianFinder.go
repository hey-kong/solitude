package main

import "container/heap"

// Leetcode 295. (hard)
type MedianFinder struct {
	h1 MaxHeap
	h2 MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		MaxHeap{},
		MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.h1.Len() == 0 || this.h1[0] >= num {
		heap.Push(&this.h1, num)
	} else {
		heap.Push(&this.h2, num)
	}
	// adjust
	if this.h1.Len() > this.h2.Len()+1 {
		x := heap.Pop(&this.h1)
		heap.Push(&this.h2, x)
	} else if this.h1.Len() < this.h2.Len() {
		x := heap.Pop(&this.h2)
		heap.Push(&this.h1, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.h1.Len() != this.h2.Len() {
		return float64(this.h1[0])
	}
	return float64(this.h1[0]+this.h2[0]) / 2
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
