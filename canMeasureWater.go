package main

// Leetcode 365. (medium)
func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	if z == 0 {
		return true
	}

	if gcd(x, y) != 0 && z%gcd(x, y) == 0 {
		return true
	}
	return false
}
