package main

// Leetcode 1248. (medium)
func numberOfSubarrays(nums []int, k int) int {
	res := 0
	memo := []int{-1}
	i := 1
	for j := 0; j <= len(nums); j++ {
		if j == len(nums) || nums[j]%2 == 1 {
			memo = append(memo, j)
		}
		if len(memo)-1-i == k {
			left := memo[i] - memo[i-1]
			right := j - memo[len(memo)-2]
			res += left * right
			i++
		}
	}
	return res
}
