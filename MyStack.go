package main

// Leetcode 225. (easy)
type MyStack struct {
	q *queueOf225
}

type queueOf225 struct {
	data []int
	size int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	q := &queueOf225{[]int{}, 0}
	return MyStack{q}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	// enqueue
	this.q.data = append(this.q.data, x)
	this.q.size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	for i := 0; i < this.q.size-1; i++ {
		// enqueue
		this.q.data = append(this.q.data, this.q.data[0])
		// dequeue
		this.q.data = this.q.data[1:]
	}
	e := this.q.data[0]
	// dequeue
	this.q.data = this.q.data[1:]
	this.q.size--
	return e
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.q.data[this.q.size-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
