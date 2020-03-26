package main

// Leetcode 914. (easy)
func hasGroupsSizeX(deck []int) bool {
	cnt := make([]int, 10000)
	for _, num := range deck {
		cnt[num]++
	}
	g := -1
	for _, num := range cnt {
		if num == 0 {
			continue
		}
		if g == -1 {
			g = num
		} else {
			g = gcd(g, num)
		}
	}
	return g >= 2
}
