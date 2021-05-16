package main

// Leetcode 5759. (easy)
func subsetXORSum(nums []int) (res int) {
	num := 1<<len(nums) - 1
	for i := 1; i <= num; i++ {
		tmpXOR := 0
		mask := 1
		for j := 0; j < len(nums); j++ {
			if mask&i != 0 {
				tmpXOR ^= nums[j]
			}
			mask <<= 1
		}
		res += tmpXOR
	}
	return
}
