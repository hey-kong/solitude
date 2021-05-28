package main

// Leetcode 477. (medium)
func totalHammingDistance(nums []int) (res int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		cnt := 0
		for _, num := range nums {
			cnt += num >> i & 1
		}
		res += cnt * (n - cnt)
	}
	return
}
