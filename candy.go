package main

// Leetcode 135. (hard)
func candy(ratings []int) int {
	lcandy := make([]int, len(ratings))
	rcandy := make([]int, len(ratings))

	lcandy[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			lcandy[i] = lcandy[i-1] + 1
		} else {
			lcandy[i] = 1
		}
	}
	rcandy[len(rcandy)-1] = 1
	for i := len(rcandy) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rcandy[i] = rcandy[i+1] + 1
		} else {
			rcandy[i] = 1
		}
	}
	res := 0
	for i := 0; i < len(ratings); i++ {
		res += max(lcandy[i], rcandy[i])
	}
	return res
}
