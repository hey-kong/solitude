package main

import "container/heap"

// Leetcode 295. (hard)
type MedianFinder struct {
	big MaxHeap
	small MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		MaxHeap{},
		MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.big.Len() == 0 || this.big[0] >= num {
		heap.Push(&this.big, num)
	} else {
		heap.Push(&this.small, num)
	}
	// adjust
	if this.big.Len() > this.small.Len()+1 {
		x := heap.Pop(&this.big)
		heap.Push(&this.small, x)
	} else if this.big.Len() < this.small.Len() {
		x := heap.Pop(&this.small)
		heap.Push(&this.big, x)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.big.Len() != this.small.Len() {
		return float64(this.big[0])
	}
	return float64(this.big[0]+this.small[0]) / 2
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
