package main

import "strconv"

// Leetcode 5789. (medium)
func numberOfRounds(startTime string, finishTime string) int {
	h, _ := strconv.Atoi(startTime[0:2])
	m, _ := strconv.Atoi(startTime[3:5])
	t0 := 60*h + m
	h, _ = strconv.Atoi(finishTime[0:2])
	m, _ = strconv.Atoi(finishTime[3:5])
	t1 := 60*h + m
	if t0 > t1 {
		t1 += 24 * 60
	}
	t1 = t1 / 15 * 15
	return max(0, (t1-t0)/15)
}
