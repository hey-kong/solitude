package main

// Leetcode 365. (medium)
func canMeasureWater(x int, y int, z int) bool {
	if z > x+y {
		return false
	}
	if z == 0 || z == x || z == y {
		return true
	}
	// ax + by = z (贝祖定理)
	if gcd(x, y) != 0 && z%gcd(x, y) == 0 {
		return true
	}
	return false
}
