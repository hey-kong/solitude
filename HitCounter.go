package main

// Leetcode 362. (medium)
type HitCounter struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		m: make(map[int]int),
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.m[timestamp]++
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	res := 0
	for k, v := range this.m {
		if k+300 <= timestamp {
			delete(this.m, k)
		} else {
			res += v
		}
	}
	return res
}
