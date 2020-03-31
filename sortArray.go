package main

// Leetcode 912. (medium)
func sortArray1(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	j := quickSortPartition(nums, left, right)
	quickSort(nums, left, j-1)
	quickSort(nums, j+1, right)
}

func quickSortPartition(nums []int, left, right int) int {
	pivot := nums[left]
	mark := left + 1
	for i := mark; i <= right; i++ {
		if nums[i] < pivot {
			nums[mark], nums[i] = nums[i], nums[mark]
			mark++
		}
	}
	nums[left], nums[mark-1] = nums[mark-1], nums[left]
	return mark - 1
}

func sortArray2(nums []int) []int {
	aux := make([]int, len(nums))
	mergeSort(nums, aux, 0, len(nums)-1)
	return nums
}

func mergeSort(nums, aux []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(nums, aux, left, mid)
	mergeSort(nums, aux, mid+1, right)
	mergeArray(nums, aux, left, mid, right)
}

func mergeArray(nums, aux []int, left, mid, right int) {
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		aux[k] = nums[k]
	}
	for k := left; k <= right; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > right {
			nums[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}
