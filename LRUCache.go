package main

// Leetcode 146. (medium)
type LRUCache struct {
	m		map[int]int
	l	 	*linkedList
	size 	int
}


func Constructor(capacity int) LRUCache {
	lru := new(LRUCache)
	lru.init(capacity)
	return *lru
}


func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	n := this.l.seek(key)
	this.l.delete(n)
	this.l.add(n)
	return this.m[key]
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.m[key]; ok {
		this.m[key] = value
		n := this.l.seek(key)
		this.l.delete(n)
		this.l.add(n)
		return
	}

	if len(this.m) == this.size {
		n := this.l.pop()
		delete(this.m, n.key)
	}
	n := &node{}
	n.key = key
	this.l.add(n)
	this.m[key] = value
}


type node struct {
	key  int
	pre  *node
	next *node
}

type linkedList struct {
	head *node
	rear *node
}

func (l *linkedList) add(n *node) {
	next := l.head.next
	l.head.next = n
	n.pre = l.head
	n.next = next
	next.pre = n
}

func (l *linkedList) seek(key int) *node {
	cur := l.head.next
	for cur != l.rear {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (l *linkedList) delete(n *node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (l *linkedList) pop() *node {
	rear := l.rear
	if rear.pre == l.head {
		return nil
	}

	n := l.rear.pre
	l.rear.pre = n.pre
	n.pre.next = l.rear
	return n
}

func (c *LRUCache) init(size int) {
	c.m = make(map[int]int, size)
	head, rear := &node{}, &node{}
	head.next = rear
	rear.next = head
	c.l = &linkedList{head, rear}
	c.size = size
}
