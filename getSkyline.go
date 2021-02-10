package main

import (
	"container/heap"
	"sort"
)

type building struct {
	right  int
	height int
}

// Leetcode 218. (hard)
func getSkyline(buildings [][]int) (res [][]int) {
	points := make([]int, 0, len(buildings)*2)
	for i := range buildings {
		points = append(points, buildings[i][0])
		points = append(points, buildings[i][1])
	}
	sort.Ints(points)

	h := buildingHeap{}
	heap.Init(&h)
	heap.Push(&h, building{right: 1<<63 - 1, height: 0})
	i := 0
	for _, p := range points {
		for i < len(buildings) && buildings[i][0] == p {
			heap.Push(&h, building{right: buildings[i][1], height: buildings[i][2]})
			i++
		}
		// 遍历横坐标时，关键点只与堆中最大高度相关，
		// 因此需要将右边缘小于等于当前点的building出队。
		// 向上的关键点：左边缘横坐标对应的高度大于堆中最大高度时出现，
		// 关键点的高度等于当前高度；
		// 向下的关键点：右边缘横坐标对应的高度等于堆中最大高度时出现，
		// 关键点的高度等于将右边缘等于当前点的堆顶building出队后的堆顶building高度。
		for h[0].right <= p {
			heap.Pop(&h)
		}
		if len(res) == 0 || res[len(res)-1][1] != h[0].height {
			res = append(res, []int{p, h[0].height})
		}
	}
	return res
}

type buildingHeap []building

func (h buildingHeap) Len() int {
	return len(h)
}

func (h buildingHeap) Less(i, j int) bool {
	return h[i].height > h[j].height
}

func (h buildingHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *buildingHeap) Push(x interface{}) {
	(*h) = append((*h), x.(building))
}

func (h *buildingHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return tmp
}
