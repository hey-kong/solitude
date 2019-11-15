package main

// Leetcode 136. (easy)
func singleNumber1(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

// Leetcode 137. (medium)
func singleNumber2(nums []int) int {
	ones := 0
	twos := 0
	threes := 0
	for _, num := range nums {
		twos |= ones & num
		ones ^= num
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}
	return ones
}

// Leetcode 260. (medium)
func singleNumber3(nums []int) []int {
	res := make([]int, 2)
	xorRes := 0
	for _, num := range nums {
		xorRes ^= num
	}

	separator := 1
	for xorRes & separator == 0 {
		separator = separator << 1
	}
	for _, num := range nums {
		if num & separator == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}