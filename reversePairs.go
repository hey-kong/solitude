package main

// Leetcode 493. (hard)
func reversePairs(nums []int) int {
	aux := make([]int, len(nums))
	return mergeSortOfReversePairs(nums, aux, 0, len(nums)-1)
}

func mergeSortOfReversePairs(nums, aux []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	ans1 := mergeSortOfReversePairs(nums, aux, left, mid)
	ans2 := mergeSortOfReversePairs(nums, aux, mid+1, right)
	i, j := left, mid+1
	ans3 := 0
	for i <= mid {
		for j <= right && nums[i] > 2*nums[j] {
			ans3 += mid - i + 1
			j++
		}
		i++
	}
	mergeArray(nums, aux, left, mid, right)
	return ans1 + ans2 + ans3
}

// Leetcode m51. (hard)
func reversePairs2(nums []int) int {
	aux := make([]int, len(nums))
	return mergeSortOfReversePairs2(nums, aux, 0, len(nums)-1)
}

func mergeSortOfReversePairs2(nums, aux []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2
	ans1 := mergeSortOfReversePairs2(nums, aux, left, mid)
	ans2 := mergeSortOfReversePairs2(nums, aux, mid+1, right)
	i, j := left, mid+1
	ans3 := 0
	for i <= mid {
		for j <= right && nums[i] > nums[j] {
			ans3 += mid - i + 1
			j++
		}
		i++
	}
	mergeArray(nums, aux, left, mid, right)
	return ans1 + ans2 + ans3
}
