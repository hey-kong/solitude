package main

import "container/list"

// Leetcode 705. (easy)
type MyHashSet struct {
	data []list.List
}

const MyHashSetBase = 769

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{data: make([]list.List, MyHashSetBase)}
}

func (this *MyHashSet) hash(key int) int {
	return key % MyHashSetBase
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[h].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}
