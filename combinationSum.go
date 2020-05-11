package main

import "sort"

// Leetcode 39. (medium)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return recursiveCombinationSum(candidates, target, 0, 0, []int{}, [][]int{})
}

func recursiveCombinationSum(candidates []int, target, cur, begin int, tmp []int, res [][]int) [][]int {
	if cur == target {
		arr := make([]int, len(tmp))
		copy(arr, tmp)
		res = append(res, arr)
		return res
	}

	for i := begin; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		cur += candidates[i]
		if cur > target {
			return res
		}
		res = recursiveCombinationSum(candidates, target, cur, i, tmp, res)
		tmp = tmp[:len(tmp)-1]
		cur -= candidates[i]
	}
	return res
}

// Leetcode 40. (medium)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return recursiveCombinationSum2(candidates, target, 0, 0, []int{}, [][]int{})
}

func recursiveCombinationSum2(candidates []int, target, cur, begin int, tmp []int, res [][]int) [][]int {
	if cur == target {
		arr := make([]int, len(tmp))
		copy(arr, tmp)
		res = append(res, arr)
		return res
	}

	for i := begin; i < len(candidates); i++ {
		if i > begin && candidates[i-1] == candidates[i] {
			continue
		}

		tmp = append(tmp, candidates[i])
		cur += candidates[i]
		if cur > target {
			return res
		}
		res = recursiveCombinationSum2(candidates, target, cur, i+1, tmp, res)
		tmp = tmp[:len(tmp)-1]
		cur -= candidates[i]
	}
	return res
}
