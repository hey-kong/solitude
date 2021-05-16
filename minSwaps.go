package main

// Leetcode 5760. (medium)
func minSwaps(s string) int {
	zeros, ones := 0, 0
	for i := range s {
		if s[i] == '0' {
			zeros++
		} else {
			ones++
		}
	}

	if len(s)%2 == 0 && zeros != ones {
		return -1
	}
	if len(s)%2 == 1 && !(zeros == ones+1 || zeros == ones-1) {
		return -1
	}

	str1, str2 := genMinSwapsStr(zeros, ones)
	changes1, changes2 := len(s), len(s)
	for i := range s {
		if s[i] == str1[i] {
			changes1--
		}
		if s[i] == str2[i] {
			changes2--
		}
	}
	return min(changes1, changes2) / 2
}

func genMinSwapsStr(zeros, ones int) (string, string) {
	str1, str2 := make([]byte, zeros+ones), make([]byte, zeros+ones)
	if zeros == ones {
		str1[0] = '0'
		str2[0] = '1'
	} else if zeros > ones {
		str1[0] = '0'
		str2[0] = '0'
	} else {
		str1[0] = '1'
		str2[0] = '1'
	}
	for i := 1; i < zeros+ones; i++ {
		if str1[i-1] == '0' {
			str1[i] = '1'
		} else {
			str1[i] = '0'
		}
		if str2[i-1] == '1' {
			str2[i] = '0'
		} else {
			str2[i] = '1'
		}
	}
	return string(str1), string(str2)
}
