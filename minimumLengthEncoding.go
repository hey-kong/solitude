package main

// Leetcode 820. (medium)
func minimumLengthEncoding(words []string) int {
	t := ConstructorOf820()
	res := 0
	for _, word := range words {
		if !t.SearchOf820(word) {
			res += t.InsertOf820(word)
		}
	}
	return res
}

func ConstructorOf820() Trie {
	return Trie{
		node:  [26]*Trie{},
		isEnd: false,
	}
}

func (this *Trie) InsertOf820(word string) int {
	l := len(word) + 1
	for i := len(word) - 1; i >= 0; i-- {
		j := word[i] - 'a'
		if this.node[j] == nil {
			this.node[j] = new(Trie)
		}
		this = this.node[j]
		if this.isEnd == true {
			this.isEnd = false
			l = i
		}
	}
	this.isEnd = true
	return l
}

func (this *Trie) SearchOf820(word string) bool {
	for i := len(word) - 1; i >= 0; i-- {
		j := word[i] - 'a'
		if this.node[j] == nil {
			return false
		}
		this = this.node[j]
	}
	return true
}
