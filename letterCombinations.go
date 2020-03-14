package main

// Leetcode 17. (medium)
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := make(map[byte]string, 8)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"
	res := dfsLetterCombinations(digits, "", 0, m, []string{})
	return res
}

func dfsLetterCombinations(digits, tmp string, idx int, m map[byte]string, res []string) []string {
	if idx == len(digits) {
		res = append(res, tmp)
		return res
	}
	str := m[digits[idx]]
	for i := range str {
		tmp += str[i : i+1]
		res = dfsLetterCombinations(digits, tmp, idx+1, m, res)
		tmp = tmp[:len(tmp)-1]
	}
	return res
}
