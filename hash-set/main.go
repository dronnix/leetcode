package main

import "fmt"

const hashSeed = 11111

func hash(key int) int {
	return key % hashSeed
}

type MyHashSet struct {
	bitmap []bool
}

func Constructor() MyHashSet {
	return MyHashSet{bitmap: make([]bool, hashSeed)}
}

func (this *MyHashSet) Add(key int) {
	this.bitmap[hash(key)] = true
}

func (this *MyHashSet) Remove(key int) {
	this.bitmap[hash(key)] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.bitmap[hash(key)]
}

func main() {
	h := Constructor()
	h.Add(42)
	fmt.Println(h.Contains(42))
}
