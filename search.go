package main

// Leetcode 33. (medium)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == nums[right] {
			break
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if target >= nums[right] && target <= nums[len(nums)-1] {
		left, right = right, len(nums)-1
	} else {
		left, right = 0, right-1
	}
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
