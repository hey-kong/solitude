package main

// Leetcode 852. (easy)
func peakIndexInMountainArray(A []int) int {
	left, right := 0, len(A)-1
	for left <= right {
		mid := left + (right-left)/2
		if A[mid] <= A[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
