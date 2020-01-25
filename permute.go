package main

// Leetcode 46. (medium)
func permute(nums []int) [][]int {
	return recursivePermute(nums, []int{}, [][]int{})
}

func recursivePermute(nums, arr []int, res [][]int) [][]int {
	if len(nums) == len(arr) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if contains(arr, nums[i]) {
			continue
		}
		arr = append(arr, nums[i])
		res = recursivePermute(nums, arr, res)
		arr = arr[:len(arr)-1]
	}
	return res
}

func contains(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}
