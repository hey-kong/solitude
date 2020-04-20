package main

// Leetcode 887. (hard)
func superEggDrop(K int, N int) int {
	memo := make([][]int, K+1)
	for i := range memo {
		memo[i] = make([]int, N+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return recursiveSuperEggDrop(K, N, memo)
}

func recursiveSuperEggDrop(K int, N int, memo [][]int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if memo[K][N] != -1 {
		return memo[K][N]
	}

	res := 10000
	lo, hi := 1, N
	for lo <= hi {
		mid := lo + (hi-lo)/2
		broken := recursiveSuperEggDrop(K-1, mid-1, memo)
		notBroken := recursiveSuperEggDrop(K, N-mid, memo)
		if broken > notBroken {
			hi = mid - 1
			res = min(res, broken+1)
		} else {
			lo = mid + 1
			res = min(res, notBroken+1)
		}
	}
	memo[K][N] = res
	return res
}
