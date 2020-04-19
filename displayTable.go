package main

import (
	"sort"
	"strconv"
)

// Leetcode 5389. (medium)
func displayTable(orders [][]string) [][]string {
	tableMap := make(map[string]int)
	foodMap := make(map[string]int)
	for i := range orders {
		if _, ok := tableMap[orders[i][1]]; !ok {
			tableMap[orders[i][1]] = -1
		}
		if _, ok := foodMap[orders[i][2]]; !ok {
			foodMap[orders[i][2]] = -1
		}
	}

	tables := []string{}
	for k := range tableMap {
		tables = append(tables, k)
	}
	sort.Sort(Table(tables))
	for i, k := range tables {
		tableMap[k] = i
	}

	foods := []string{}
	for k := range foodMap {
		foods = append(foods, k)
	}
	sort.Sort(sort.StringSlice(foods))
	for i, k := range foods {
		foodMap[k] = i
	}

	memo := make([][]int, len(tableMap))
	for j := range memo {
		memo[j] = make([]int, len(foodMap))
	}
	for j := range orders {
		memo[tableMap[orders[j][1]]][foodMap[orders[j][2]]]++
	}
	res := [][]string{}
	res = append(res, append([]string{"Table"}, foods...))
	for i := range tables {
		tmp := []string{}
		tmp = append(tmp, tables[i])
		for j := range memo[i] {
			tmp = append(tmp, strconv.Itoa(memo[i][j]))
		}
		res = append(res, tmp)
	}
	return res
}

type Table []string

func (arr Table) Len() int {
	return len(arr)
}

func (arr Table) Less(i, j int) bool {
	a, _ := strconv.Atoi(arr[i])
	b, _ := strconv.Atoi(arr[j])
	return a < b
}

func (arr Table) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
