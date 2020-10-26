package main

// Leetcode m16.10. (medium)
func maxAliveYear(birth []int, death []int) int {
	arr := make([]int, 102)
	for i := 0; i < len(birth); i++ {
		arr[birth[i]-1900]++
		arr[death[i]-1900+1]--
	}
	alive := 0
	maxAlive := 0
	res := 0
	for i := 0; i < 101; i++ {
		alive += arr[i]
		if alive > maxAlive {
			maxAlive = alive
			res = i + 1900
		}
	}
	return res
}
