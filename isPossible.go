package main

// Leetcode 659. (medium)
func isPossible(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	endCnt := make(map[int]int)
	for _, num := range nums {
		if m[num] == 0 {
			continue
		}
		if endCnt[num-1] > 0 {
			endCnt[num-1]--
			endCnt[num]++
			m[num]--
		} else if m[num+1] > 0 && m[num+2] > 0 {
			m[num]--
			m[num+1]--
			m[num+2]--
			endCnt[num+2]++
		} else {
			return false
		}
	}
	return true
}
