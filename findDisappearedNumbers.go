package main

// Leetcode 448. (easy)
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	res := []int{}
	for i, num := range nums {
		if num != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
