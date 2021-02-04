package main

import (
	"container/heap"
)

// Leetcode 480. (hard)
func medianSlidingWindow(nums []int, k int) (res []float64) {
	m := MedianFinder{
		MaxHeap{},
		MinHeap{},
	}
	for i := 0; i < k; i++ {
		heap.Push(&m.big, nums[i])
	}
	for i := 0; i < k/2; i++ {
		heap.Push(&m.small, heap.Pop(&m.big))
	}
	res = append(res, m.findSlidingWindowMedian(k))

	tbd := map[int]int{}
	for i := k; i < len(nums); i++ {
		diff := 0
		if nums[i-k] <= m.big[0] {
			diff++
		} else {
			diff--
		}
		tbd[nums[i-k]]++
		if nums[i] <= m.big[0] {
			heap.Push(&m.big, nums[i])
			diff--
		} else {
			heap.Push(&m.small, nums[i])
			diff++
		}

		if diff > 0 {
			heap.Push(&m.big, heap.Pop(&m.small))
		}
		if diff < 0 {
			heap.Push(&m.small, heap.Pop(&m.big))
		}

		for m.big.Len() != 0 && tbd[m.big[0]] != 0 {
			tbd[m.big[0]]--
			heap.Pop(&m.big)
		}
		for m.small.Len() != 0 && tbd[m.small[0]] != 0 {
			tbd[m.small[0]]--
			heap.Pop(&m.small)
		}
		res = append(res, m.findSlidingWindowMedian(k))
	}
	return
}

func (m *MedianFinder) findSlidingWindowMedian(k int) float64 {
	if k%2 == 1 {
		return float64(m.big[0])
	}
	return float64(m.big[0]+m.small[0]) / 2
}
