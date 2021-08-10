package main

// Leetcode 5. (medium)
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start := 0
	end := 0
	for i := range s {
		l1, r1 := findPalindrome(s, i-1, i+1)
		l2, r2 := findPalindrome(s, i, i+1)
		if (r1-l1 > r2-l2) && (r1-l1 > end-start) {
			start = l1
			end = r1
		} else if (r2-l2 > r1-l1) && (r2-l2 > end-start) {
			start = l2
			end = r2
		}
	}
	return s[start : end+1]
}

func findPalindrome(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}

func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	start := 0
	end := 0
	for l := 2; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			if s[i] == s[i+l-1] && (l <= 3 || dp[i+1][i+l-2]) {
				dp[i][i+l-1] = true
				start, end = i, i+l-1
			}
		}
	}
	return s[start : end+1]
}

// Leetcode 409. (easy)
func longestPalindrome3(s string) int {
	charCnt := make([]int, 58)
	for _, char := range s {
		charCnt[char-'A']++
	}
	res := 0
	resIsOdd := 0
	for _, num := range charCnt {
		if num%2 == 0 {
			res += num
		} else {
			res += num / 2 * 2
			resIsOdd = 1
		}
	}
	return res + resIsOdd
}
