package main

// Leetcode 1567. (medium)
func getMaxLen(nums []int) int {
	pos, neg := 0, 0
	res := 0
	if nums[0] < 0 {
		neg = 1
	} else if nums[0] > 0 {
		pos = 1
		res = 1
	}
	for _, num := range nums[1:] {
		if num < 0 {
			prePos, preNeg := pos, neg
			if preNeg == 0 {
				pos = 0
			} else {
				pos = preNeg + 1
			}
			neg = prePos + 1
		} else if num > 0 {
			pos++
			if neg == 0 {
				neg = 0
			} else {
				neg++
			}
		} else {
			pos = 0
			neg = 0
		}
		res = max(res, pos)
	}
	return res
}
