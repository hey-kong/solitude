package main

import "sort"

// Leetcode LCP.18. (easy)
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	i, j := 0, len(drinks)-1
	for i < len(staple) && j >= 0 {
		if staple[i]+drinks[j] <= x {
			res += j + 1
			i++
		} else {
			j--
		}
	}
	return res % 1000000007
}
