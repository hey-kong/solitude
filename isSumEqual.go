package main

import "strconv"

// Leetcode 5772. (easy)
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	m := make(map[rune]int)
	r := 'a'
	for i := 0; i < 26; i++ {
		m[r] = i
		r++
	}

	buf := ""
	for _, r := range firstWord {
		buf += strconv.Itoa(m[r])
	}
	firstNum, _ := strconv.Atoi(buf)
	buf = ""
	for _, r := range secondWord {
		buf += strconv.Itoa(m[r])
	}
	secondNum, _ := strconv.Atoi(buf)
	buf = ""
	for _, r := range targetWord {
		buf += strconv.Itoa(m[r])
	}
	targetNum, _ := strconv.Atoi(buf)
	return firstNum+secondNum == targetNum
}
