package main

// Leetcode 47. (medium)
func permuteUnique(nums []int) [][]int {
    return recursivePermuteUnique(nums, []int{}, [][]int{})
}

func recursivePermuteUnique(nums, arr []int, res [][]int) [][]int {
    if len(nums) == len(arr) {
		tmp := make([]int, len(arr))
		for i, j := range arr {
			tmp[i] = nums[j]
		}
		res = append(res, tmp)
		return res
	}

	used := []int{}
	for i := 0; i < len(nums); i++ {
		if contains(used, nums[i]) || contains(arr, i) {
			continue
		}
		arr = append(arr, i)
		res = recursivePermuteUnique(nums, arr, res)
		arr = arr[:len(arr)-1]
		used = append(used, nums[i])
	}
	return res
}
