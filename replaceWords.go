package main

import "strings"

// Leetcode 648. (medium)
func replaceWords(dict []string, sentence string) string {
	t := &Trie{}
	for _, word := range dict {
		t.Insert(word)
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = t.findRoot(word)
	}
	return strings.Join(words, " ")
}

func (t *Trie) findRoot(word string) string {
	for j, c := range word {
		i := c - 'a'
		if t.node[i] == nil {
			break
		}
		t = t.node[i]
		if t.isEnd {
			return word[:j+1]
		}
	}
	return word
}
