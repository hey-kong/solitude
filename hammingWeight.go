package main

// Leetcode 191. (easy)
func hammingWeight(num uint32) int {
	bits := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num & mask != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}
