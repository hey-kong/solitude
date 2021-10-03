package main

// Leetcode 5891. (medium)
func missingRolls(rolls []int, mean int, n int) (res []int) {
	sum := 0
	for _, num := range rolls {
		sum += num
	}
	sum = (len(rolls)+n)*mean - sum
	if sum < n || sum > 6*n {
		return
	}

	redundancy := sum % n
	val := sum / n
	i := 0
	for i < redundancy {
		res = append(res, val+1)
		i++
	}
	for i < n {
		res = append(res, val)
		i++
	}
	return
}
