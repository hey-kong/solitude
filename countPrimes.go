package main

// Leetcode 204. (easy)
func countPrimes(n int) int {
	m := make([]bool, n)
	for i := range m {
		m[i] = true
	}

	for i := 2; i * i < n; i++ {
		if m[i] {
			for j := i * i; j < n; j += i {
				m[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if m[i] {
			count++
		}
	}
	return count
}
