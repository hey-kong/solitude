package main

// Leetcode 208. (medium)
type Trie struct {
	node  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node:  [26]*Trie{},
		isEnd: false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		i := c - 'a'
		if this.node[i] == nil {
			this.node[i] = new(Trie)
		}
		this = this.node[i]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		i := c - 'a'
		if this.node[i] == nil {
			return false
		}
		this = this.node[i]
	}
	return this.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		i := c - 'a'
		if this.node[i] == nil {
			return false
		}
		this = this.node[i]
	}
	return true
}
