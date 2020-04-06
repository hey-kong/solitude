package main

// Leetcode 460. (hard)
type LFUCache struct {
	qm  map[int][3]int // key => val, times, time
	t   int            // time
	cap int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		qm:  make(map[int][3]int),
		t:   0,
		cap: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if !this.exist(key) {
		return -1
	}
	this.t++
	this.qm[key] = [3]int{this.qm[key][0], this.qm[key][1] + 1, this.t}
	return this.qm[key][0]
}

func (this *LFUCache) Put(key int, value int) {
	exist := this.exist(key)
	this.t++

	if !exist {
		if len(this.qm) != this.cap {
			this.qm[key] = [3]int{value, 1, this.t}
		} else {
			lastKey, i := 0, 0
			for k := range this.qm {
				if i == 0 {
					lastKey = k
					i++
					continue
				}
				if this.qm[lastKey][1] > this.qm[k][1] || (this.qm[lastKey][1] == this.qm[k][1] && this.qm[lastKey][2] > this.qm[k][2]) {
					lastKey = k
				}
			}
			this.qm[key] = [3]int{value, 1, this.t}
			delete(this.qm, lastKey)
		}
	}

	if exist {
		this.qm[key] = [3]int{value, this.qm[key][1] + 1, this.t}
	}
}

func (this *LFUCache) exist(key int) bool {
	_, e := this.qm[key]
	return e
}
