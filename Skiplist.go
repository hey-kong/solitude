package main

import (
	"math"
	"math/rand"
)

// Leetcode 1206. (hard)
type Skiplist struct {
	head  *skiplistNode
	heads []*skiplistNode
}

var MAX_LEVEL int = 16

type skiplistNode struct {
	val         int
	right, down *skiplistNode
}

func Constructor() Skiplist {
	list := Skiplist{
		head:  &skiplistNode{},
		heads: []*skiplistNode{},
	}
	list.heads = append(list.heads, list.head)
	return list
}

func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for cur != nil {
		if cur.right != nil {
			if cur.right.val < target {
				cur = cur.right
			} else if cur.right.val > target {
				cur = cur.down
			} else {
				return true
			}
		} else {
			cur = cur.down
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	cur := this.head
	for {
		if cur.right != nil {
			if cur.right.val <= num {
				cur = cur.right
			} else if cur.down != nil {
				cur = cur.down
			} else {
				break
			}
		} else {
			if cur.down != nil {
				cur = cur.down
			} else {
				break
			}
		}
	}
	cur.right = &skiplistNode{
		val:   num,
		right: cur.right,
		down:  nil,
	}
	if len(this.heads) < MAX_LEVEL && rand.Float64() < math.Pow(0.5, float64(len(this.heads))) {
		up := this.heads[len(this.heads)-1]
		downHead := &skiplistNode{}
		down := downHead
		for up != nil {
			up.down = down
			if up.right != nil {
				down.right = &skiplistNode{
					val: up.right.val,
				}
			}
			up = up.right
			down = down.right
		}
		this.heads = append(this.heads, downHead)
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur := this.head
	find := false
	for cur != nil {
		if cur.right != nil {
			if cur.right.val < num {
				cur = cur.right
			} else if cur.right.val > num {
				cur = cur.down
			} else {
				find = true
				cur.right = cur.right.right
				cur = cur.down
			}
		} else {
			cur = cur.down
		}
	}
	return find
}
