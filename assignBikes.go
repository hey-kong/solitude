package main

import (
	"math"
	"sort"
)

type WorkerBike struct {
	dis		int
	wid		int
	bid		int
}

type WorkerBikeList []WorkerBike

// Leetcode 1057. (medium)
func assignBikes(workers [][]int, bikes [][]int) []int {
	m, n := len(workers), len(bikes)
	wMatched := make([]bool, m)
	bMatched := make([]bool, n)
	list := make(WorkerBikeList, m*n)
	res := make([]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			list[i*n+j] = WorkerBike{
				dis: distanceManhattan(workers[i][0],workers[i][1],bikes[j][0],bikes[j][1]),
				wid: i,
				bid: j,
			}
		}
	}

	sort.Sort(list)
	cnt := 0
	for _, tmp := range list {
		if !wMatched[tmp.wid] && !bMatched[tmp.bid] {
			res[tmp.wid] = tmp.bid
			cnt++
			wMatched[tmp.wid] = true
			bMatched[tmp.bid] = true
		}
		if cnt == n {
			break
		}
	}
	return res
}

func distanceManhattan(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func (list WorkerBikeList) Len() int {
	return len(list)
}

func (list WorkerBikeList) Less(i int, j int) bool {
	if list[i].dis == list[j].dis && list[i].wid == list[j].wid {
		return list[i].bid < list[j].bid
	}
	if list[i].dis == list[j].dis {
		return list[i].wid < list[j].wid
	}
	return list[i].dis < list[j].dis
}

func (list WorkerBikeList) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}
