package main

// Leetcode 151. (medium)
func reverseWords(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}

	i := 0
	start := 0
	for {
		for start < len(b) && b[start] == 32 {
			start++
		}
		if start == len(b) {
			break
		}
		end := start
		for end < len(b) && b[end] != 32 {
			end++
		}
		for j := start; j < (start + end) / 2; j++ {
			b[j], b[start+end-1-j] = b[start+end-1-j], b[j]
		}

		if i != 0 {
			b[i] = 32
			i++
		}
		for start < end {
			b[i] = b[start]
			i++
			start++
		}
		start = end
	}
	return string(b[:i])
}

func reverseWords2(s []byte)  {
	for i := 0 ; i < len(s) / 2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	start := 0
	for start < len(s){
		end := start
		for end < len(s) && s[end] != 32 {
			end++
		}

		for i := start; i < (start + end) / 2; i++ {
			s[i], s[start+end-1-i] = s[start+end-1-i], s[i]
		}
		start = end + 1
	}
}
