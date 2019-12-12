package main

import "strconv"

// Leetcode 412. (easy)
func fizzBuzz(n int) []string {
	res := []string{}

	var s string
	for i := 1; i <= n; i++ {
		s = ""
		if i % 3 == 0 {
			s += "Fizz"
		}
		if i % 5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		res = append(res, s)
	}

	return res
}
