package main

import "fmt"

// The task: https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	root *node
}

type node struct {
	symbols   map[byte]*node
	terminate bool
}

func newNode() *node {
	return &node{symbols: make(map[byte]*node)}
}

func (n *node) insert(s []byte) {
	if _, ok := n.symbols[s[0]]; !ok {
		n.symbols[s[0]] = newNode()
	}
	if len(s) == 1 {
		n.symbols[s[0]].terminate = true
		return
	}
	n.symbols[s[0]].insert(s[1:])
}

func (n *node) search(s []byte, prefix bool) bool {
	nn, ok := n.symbols[s[0]]
	if !ok {
		return false
	}
	if len(s) == 1 {
		if prefix {
			return true
		}
		return nn.terminate
	}
	return nn.search(s[1:], prefix)
}

func Constructor() Trie {
	return Trie{root: newNode()}
}

func (t *Trie) Insert(word string) {
	t.root.insert([]byte(word))
}

func (t *Trie) Search(word string) bool {
	if word == "" {
		return false
	}
	return t.root.search([]byte(word), false)
}

func (t *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return false
	}
	return t.root.search([]byte(prefix), true)
}

func main() {
	trie := Constructor()
	trie.Insert("Foo")
	trie.Insert("FooBar")
	fmt.Println(trie.Search("FooBar"))
	fmt.Println(trie.Search("FooB"))
	fmt.Println(trie.StartsWith("FooB"))
}
