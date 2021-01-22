package main

// Leetcode 438. (medium)
func findAnagrams(s string, p string) (res []int) {
	ns, np := len(s), len(p)
	if ns < np {
		return
	}

	arrS, arrP := make([]int, 26), make([]int, 26)
	i := 0
	for i < np {
		arrS[s[i]-'a']++
		arrP[p[i]-'a']++
		i++
	}
	for i < ns {
		if sliceEqual(arrS, arrP) {
			res = append(res, i-np)
		}
		arrS[s[i-np]-'a']--
		arrS[s[i]-'a']++
		i++
	}
	if sliceEqual(arrS, arrP) {
		res = append(res, i-np)
	}
	return
}

func sliceEqual(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
