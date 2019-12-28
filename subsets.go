package main

// Leetcode 78. (medium)
func subsets(nums []int) [][]int {
	l := len(nums)
	res := [][]int{}
	for i := 0; i < 1 << l; i++ {
		arr := []int{}
		for k := 0; k < l; k++ {
			if (i >> k) & 1 == 1 {
				arr = append(arr, nums[k])
			}
		}
		res = append(res, arr)
	}
	return res
}
