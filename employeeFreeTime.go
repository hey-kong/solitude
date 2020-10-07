package main

import "container/heap"

// Leetcode 759. (hard)
type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	if len(schedule) == 0 || schedule == nil {
		return nil
	}

	var h MinIntervalHeap
	heap.Init(&h)
	for _, s := range schedule {
		for _, tmp := range s {
			heap.Push(&h, tmp)
		}
	}

	res := []*Interval{}
	end := h[0].Start
	for len(h) > 0 {
		tmp := heap.Pop(&h).(*Interval)
		if tmp.Start <= end {
			end = max(end, tmp.End)
		} else {
			res = append(res, &Interval{
				Start: end,
				End:   tmp.Start,
			})
			end = tmp.End
		}
	}
	return res
}

type MinIntervalHeap []*Interval

func (h MinIntervalHeap) Len() int {
	return len(h)
}

func (h MinIntervalHeap) Less(i, j int) bool {
	if h[i].Start == h[j].Start {
		return h[i].End < h[j].End
	}
	return h[i].Start < h[j].Start
}

func (h MinIntervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinIntervalHeap) Push(x interface{}) {
	(*h) = append((*h), x.(*Interval))
}

func (h *MinIntervalHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return tmp
}
