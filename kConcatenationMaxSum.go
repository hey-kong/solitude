package main

// Leetcode 1191. (medium)
func kConcatenationMaxSum(arr []int, k int) int {
	res1 := 0
	sum := 0
	for _, num := range arr {
		if sum <= 0 {
			sum = num
		} else {
			sum += num
		}
		res1 = max(res1, sum)
	}
	if k == 1 {
		return res1
	}

	n := len(arr)
	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = arr[0], arr[n-1]
	maxL, maxR := max(arr[0], 0), max(arr[n-1], 0)
	for i := 1; i < n; i++ {
		left[i] += arr[i] + left[i-1]
		maxL = max(maxL, left[i])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] += arr[i] + right[i+1]
		maxR = max(maxR, right[i])
	}
	res2 := maxL + maxR
	if right[0] >= 0 {
		res2 += (k - 2) * right[0]
	}
	return max(res1, res2) % 1000000007
}
