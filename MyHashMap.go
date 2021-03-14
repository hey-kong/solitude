package main

import "container/list"

// Leetcode 706. (easy)
type MyHashMap struct {
	data []list.List
}

type MyHashMapEntry struct {
	key, value int
}

const MyHashMapBase int = 1000

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, MyHashMapBase)}
}

func (this *MyHashMap) hash(key int) int {
	return key % MyHashMapBase
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	newEntry := MyHashMapEntry{key: key, value: value}
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		entry := e.Value.(MyHashMapEntry)
		if entry.key == key {
			e.Value = newEntry
			return
		}
	}
	this.data[h].PushBack(newEntry)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		entry := e.Value.(MyHashMapEntry)
		if entry.key == key {
			return entry.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		entry := e.Value.(MyHashMapEntry)
		if entry.key == key {
			this.data[h].Remove(e)
		}
	}
}
