package main

import "container/list"

type LRUCache struct {
	capacity int
	table    map[int]record
	queue    *list.List
}

type record struct {
	value   int
	element *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		table:    make(map[int]record, capacity),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	rec, ok := c.table[key]
	if !ok {
		return -1
	}
	c.queue.MoveToFront(rec.element)
	return rec.value
}

func (c *LRUCache) Put(key, value int) {
	newElement := c.queue.PushFront(key)
	oldRecord, ok := c.table[key]
	if ok {
		c.queue.Remove(oldRecord.element)
	}
	c.table[key] = record{element: newElement, value: value}
	if c.queue.Len() > c.capacity {
		lruRecord := c.queue.Back()
		delete(c.table, lruRecord.Value.(int))
		c.queue.Remove(lruRecord)
	}
}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	if lRUCache.Get(1) != 1 {
		panic("1")
	}
	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	if lRUCache.Get(2) != -1 {
		panic("2")
	}
	lRUCache.Put(4, 4)         // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	if lRUCache.Get(1) != -1 { // return -1 (not found)
		panic("-")
	}
	if lRUCache.Get(3) != 3 { // return 3
		panic("3")
	}
	if lRUCache.Get(4) != 4 { // return 3
		panic("4")
	}
}
