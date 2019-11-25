package main

// Leetcode 191. (easy)
func hammingWeight1(num uint32) int {
	bits := 0
	for num != 0 {
		bits++
		num = num & (num - 1)
	}
	return bits
}

func hammingWeight2(num uint32) int {
	num = (num&0x55555555) + (num>>1&0x55555555)
	num = (num&0x33333333) + (num>>2&0x33333333)
	num = (num&0x0F0F0F0F) + (num>>4&0x0F0F0F0F)
	num = (num&0x00FF00FF) + (num>>8&0x00FF00FF)
	num = (num&0x0000FFFF) + (num>>16&0x0000FFFF)
	return int(num)
}