package main

// Leetcode 169. (easy)
func majorityElement(nums []int) int {
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		} else {
			if nums[i] == res {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return res
}
