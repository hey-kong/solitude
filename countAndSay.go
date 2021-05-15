package main

// Leetcode 38. (easy)
func countAndSay(n int) string {
	res := []byte{'1'}
	for i := 1; i < n; i++ {
		tmp := []byte{}
		for j := 0; j < len(res); j++ {
			cnt := 1
			for j+1 < len(res) && res[j] == res[j+1] {
				cnt++
				j++
			}
			tmp = append(tmp, byte(cnt+'0'), res[j])
		}
		res = tmp
	}
	return string(res)
}
