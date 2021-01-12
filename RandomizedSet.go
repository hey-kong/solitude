package main

import "math/rand"

// Leetcode 380. (medium)
type RandomizedSet struct {
	arr []int
	m   map[int]int
}

/** Initialize your data strcture here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		m:   make(map[int]int, 0),
	}
}

/* Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

/* Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	i := this.m[val]
	this.arr[i], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[i]
	this.m[this.arr[i]] = i
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, val)
	return true
}

/* Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.arr))
	return this.arr[i]
}
