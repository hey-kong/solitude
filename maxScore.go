package main

// Leetcode 5392. (easy)
func maxScore(s string) int {
	res := 0
	if s[0] == '0' {
		res++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			res++
		}
	}

	tmp := res
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '1' {
			tmp--
		}
		if s[i] == '0' {
			tmp++
			res = max(res, tmp)
		}
	}
	return res
}

// Leetcode 5393. (medium)
func maxScore2(cardPoints []int, k int) int {
	res := 0
	for i := 0; i < k; i++ {
		res += cardPoints[i]
	}

	tmp := res
	for i := 0; i < k; i++ {
		tmp -= cardPoints[k-1-i]
		tmp += cardPoints[len(cardPoints)-1-i]
		res = max(res, tmp)
	}
	return res
}
