package main

// Leetcode 134. (medium)
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		sum += gas[i]
	}
	if sum < 0 {
		return -1
	}

	start := 0
	sum = 0
	for i := range gas {
		sum += gas[i]
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	return start
}
