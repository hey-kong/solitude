package main

// Leetcode 605. (easy)
func canPlaceFlowers(flowerbed []int, n int) bool {
	can := 0
	i := 0
	for i < len(flowerbed) {
		if flowerbed[i] == 0 && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			can++
			i += 2
		} else if flowerbed[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return can >= n
}
