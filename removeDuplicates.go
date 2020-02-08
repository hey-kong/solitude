package main

// Leetcode 26. (easy)
func removeDuplicates(nums []int) int {
	mark := 0
	for _, num := range nums {
		if nums[mark] != num {
			mark++
			nums[mark] = num
		}
	}
	return mark + 1
}

// Leetcode 80. (medium)
func removeDuplicates2(nums []int) int {
	mark := 0
	markIsSecond := false
	for i := 1; i < len(nums); i++ {
		if nums[mark] != nums[i] {
			mark++
			nums[mark] = nums[i]
			markIsSecond = false
		} else if !markIsSecond {
			mark++
			nums[mark] = nums[i]
			markIsSecond = true
		}
	}
	return mark + 1
}
