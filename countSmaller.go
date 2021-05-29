package main

// Leetcode 315. (hard)
func countSmaller(nums []int) []int {
	idx, aux, res := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	for i := range nums {
		idx[i] = i
	}
	return mergeSortCountSmaller(nums, idx, aux, 0, len(nums)-1, res)
}

func mergeSortCountSmaller(nums, idx, aux []int, left, right int, res []int) []int {
	if left >= right {
		return res
	}

	mid := left + (right-left)/2
	res = mergeSortCountSmaller(nums, idx, aux, left, mid, res)
	res = mergeSortCountSmaller(nums, idx, aux, mid+1, right, res)
	i, j := left, mid+1
	for j <= right {
		for i <= mid && nums[idx[i]] > nums[idx[j]] {
			res[idx[i]] += right - j + 1
			i++
		}
		j++
	}
	mergeCountSmaller(nums, idx, aux, left, mid, right)
	return res
}

func mergeCountSmaller(nums, idx, aux []int, left, mid, right int) {
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		aux[k] = idx[k]
	}
	for k := left; k <= right; k++ {
		if i > mid {
			idx[k] = aux[j]
			j++
		} else if j > right {
			idx[k] = aux[i]
			i++
		} else if nums[aux[i]] < nums[aux[j]] {
			idx[k] = aux[j]
			j++
		} else {
			idx[k] = aux[i]
			i++
		}
	}
}
