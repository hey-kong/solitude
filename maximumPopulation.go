package main

// Leetcode 5750. (easy)
func maximumPopulation(logs [][]int) int {
	offset := 1950
	arr := make([]int, 101)
	for i := range logs {
		arr[logs[i][0]-offset]++
		arr[logs[i][1]-offset]--
	}
	mx, cur, res := 0, 0, 0
	for i := range arr {
		cur += arr[i]
		if cur > mx {
			mx = cur
			res = i
		}
	}
	return res + offset
}
