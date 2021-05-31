package main

// Leetcode m17.06. (hard)
func numberOf2sInRange(n int) (res int) {
	H := n / 10
	i := n % 10
	L := 0
	base := 1
	for L < n {
		res += H * base
		if i > 2 {
			res += base
		} else if i == 2 {
			res += L + 1
		}

		L += i * base
		i = H % 10
		H /= 10
		base *= 10
	}
	return
}
