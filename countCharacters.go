package main

// Leetcode 1160. (easy)
func countCharacters(words []string, chars string) int {
	var cnt [26]int
	for _, char := range chars {
		cnt[char-'a']++
	}

	res := 0
	for _, word := range words {
		tmp := cnt
		flag := true
		for _, char := range word {
			tmp[char-'a']--
			if tmp[char-'a'] < 0 {
				flag = false
				break
			}
		}
		if flag {
			res += len(word)
		}
	}
	return res
}
