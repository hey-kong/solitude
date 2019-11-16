package main

import "sort"

// Leetcode 15. (medium)
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	res := [][]int{}
	sort.Ints(nums)
	for i, num := range nums {
		if num > 0 || i == len(nums) - 2 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}

		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			if num + nums[begin] + nums[end] == 0 {
				res = append(res, []int{num, nums[begin], nums[end]})
				for begin < end {
					begin++
					if nums[begin-1] != nums[begin] {
						break
					}
				}
				for begin < end {
					end--
					if nums[end] != nums[end+1] {
						break
					}
				}
			} else if num + nums[begin] + nums[end] < 0 {
				begin++
			} else {
				end--
			}
		}
	}
	return res
}
