package main

// Leetcode 29. (medium)
func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	if divisor == 1 {
		return dividend
	}

	same := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}
	arr := make([]int, 33)
	arr[0] = divisor
	i := 1
	for i < 33 {
		if arr[i-1] < -(-dividend >> 1) {
			break
		}
		arr[i] = -(-arr[i-1] << 1)
		i++
	}
	i--
	res := 0
	for i >= 0 {
		if dividend <= arr[i] {
			res += 1 << uint(i)
			dividend -= arr[i]
		}
		i--
	}
	if !same {
		res = -res
	}
	return res
}
