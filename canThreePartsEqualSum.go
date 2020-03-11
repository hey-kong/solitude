package main

// Leetcode 1013. (easy)
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, num := range A {
		sum += num
	}
	if sum%3 != 0 {
		return false
	}
	target := sum / 3
	i, j := 0, len(A)-1
	left, right := 0, 0
	for i < j {
		left += A[i]
		if left == target {
			break
		}
		i++
	}
	if left != target {
		return false
	}
	for i < j {
		right += A[j]
		if right == target {
			break
		}
		j--
	}
	if right != target {
		return false
	}
	if j-i > 1 {
		return true
	}
	return false
}
