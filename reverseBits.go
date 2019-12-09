package main

// Leetcode 190. (easy)
func reverseBits(num uint32) uint32 {
	num = (((num & 0xffff0000) >> 16) & 0x0000ffff) | ((num & 0x0000ffff) << 16)
	num = (((num & 0xff00ff00) >> 8) & 0x00ff00ff) | ((num & 0x00ff00ff) << 8)
	num = (((num & 0xf0f0f0f0) >> 4) & 0x0f0f0f0f) | ((num & 0x0f0f0f0f) << 4)
	num = (((num & 0xcccccccc) >> 2) & 0x33333333) | ((num & 0x33333333) << 2)
	num = (((num & 0xaaaaaaaa) >> 1) & 0x55555555) | ((num & 0x55555555) << 1)
	return num
}
