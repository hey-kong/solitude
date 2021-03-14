package main

import "reflect"

// Leetcode 5701. (easy)
func areAlmostEqual(s1 string, s2 string) bool {
	m1, m2 := make(map[byte]int), make(map[byte]int)
	diff := 0
	for i := range s1 {
		m1[s1[i]]++
		m2[s2[i]]++
		if s1[i] != s2[i] {
			diff++
		}
	}
	return (diff == 2 || diff == 0) && reflect.DeepEqual(m1, m2)
}
