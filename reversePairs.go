package main

// Leetcode m51. (hard)
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
	ans3 := mergeOfReversePairs(nums, aux, left, mid, right)
	return ans1 + ans2 + ans3
}

func mergeOfReversePairs(nums, aux []int, left, mid, right int) int {
	i, j := left, mid+1
	ans := 0
	for k := left; k <= right; k++ {
		aux[k] = nums[k]
	}
	for k := left; k <= right; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > right {
			nums[k] = aux[i]
			ans += right - mid
			i++
		} else if aux[j] < aux[i] {
			nums[k] = aux[j]
			j++
		} else {
			nums[k] = aux[i]
			ans += j - mid - 1
			i++
		}
	}
	return ans
}
