package main

// Leetcode 1175. (easy)
func numPrimeArrangements(n int) int {
	primes := countPrimes(n + 1)
	return factorialWithMod(primes) * factorialWithMod(n-primes) % 1000000007
}

func factorialWithMod(n int) int {
	res := 1
	mod := 1000000007
	for i := 1; i <= n; i++ {
		res *= i
		if res >= mod {
			res %= mod
		}
	}
	return res
}
