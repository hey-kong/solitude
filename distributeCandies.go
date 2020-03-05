package main

import "math"

// Leetcode 1103. (easy)
func distributeCandies(candies int, num_people int) []int {
	items := (int)(math.Floor(-0.5 + math.Sqrt(0.25+2*float64(candies))))
	rows := items / num_people
	remainder := items % num_people
	left := candies - (1+items)*items/2

	res := make([]int, num_people)
	for i := 0; i < num_people; i++ {
		res[i] = rows*(i+1) + rows*(rows-1)/2*num_people
		if i+1 <= remainder {
			res[i] += (i + 1) + rows*num_people
		}
	}
	res[remainder] += left
	return res
}
