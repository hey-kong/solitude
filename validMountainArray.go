package main

// Leetcode 941. (easy)
func validMountainArray(A []int) bool {
	if len(A) <= 2 {
		return false
	}

	i := 0
	for i < len(A)-1 {
		if A[i] >= A[i+1] {
			break
		}
		i++
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for i < len(A)-1 {
		if A[i] <= A[i+1] {
			break
		}
		i++
	}
	return i == len(A)-1
}
