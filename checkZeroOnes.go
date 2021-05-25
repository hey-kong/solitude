package main

// Leetcode 1869. (easy)
func checkZeroOnes(s string) bool {
	zero, one := 0, 0
	mZero, mOne := 0, 0
	for i := range s {
		if s[i] == '0' {
			zero++
			mZero = max(mZero, zero)
			one = 0
		} else {
			one++
			mOne = max(mOne, one)
			zero = 0
		}
	}
	return mZero < mOne
}
