package main

import "sort"

// Leetcode 49. (medium)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		key := getAnagramKey(str)
		m[key] = append(m[key], str)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getAnagramKey(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i int, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
