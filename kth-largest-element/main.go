package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kl := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kl.Add(3))
	fmt.Println(kl.Add(5))
	fmt.Println(kl.Add(10))
	fmt.Println(kl.Add(9))
	fmt.Println(kl.Add(4))
	// 4,5,5,8,8
}

type KthLargest struct {
	h intHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{k: k, h: nums}
	heap.Init(&result.h)
	for result.h.Len() > k {
		heap.Pop(&result.h)
	}
	return result
}

func (l *KthLargest) Add(val int) int {
	heap.Push(&l.h, val)
	if l.h.Len() > l.k {
		heap.Pop(&l.h)
	}
	return l.h[0]
}

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	l := len(*h) - 1
	r := (*h)[l]
	*h = (*h)[:l]
	return r
}
