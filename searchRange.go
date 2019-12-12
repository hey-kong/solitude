package main

// Leetcode 34. (medium)
func searchRange(nums []int, target int) []int {
	l := leftBound(nums, target)
	r := rightBound(nums, target)
	return []int{l, r}
}

func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
	}

	if left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
	}

	if right < 0 {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
	}
	return -1
}
