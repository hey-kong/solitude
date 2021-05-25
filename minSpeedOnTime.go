package main

import "math"

// Leetcode 1870. (medium)
func minSpeedOnTime(dist []int, hour float64) int {
	if hour <= math.Floor(float64(len(dist)-1)) {
		return -1
	}

	left, right := 1, 10000000
	for left <= right {
		mid := left + (right-left)/2
		tmpHour := 0.0
		for _, d := range dist[:len(dist)-1] {
			tmpHour += math.Ceil(float64(d) / float64(mid))
		}
		tmpHour += float64(float64(dist[len(dist)-1]) / float64(mid))

		if tmpHour <= hour {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
