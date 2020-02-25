package main

// Leetcode 313. (medium)
func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]int, n)
	nums[0] = 1
	idx := make([]int, len(primes))
	res := make([]int, len(primes))
	for i := 1; i < n; i++ {
		tmp := 1<<31 - 1
		for j := range primes {
			res[j] = nums[idx[j]] * primes[j]
			tmp = min(tmp, res[j])
		}
		nums[i] = tmp
		for j := range primes {
			if tmp == res[j] {
				idx[j]++
			}
		}
	}
	return nums[n-1]
}
