package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type headHeap []*ListNode

func (h *headHeap) Len() int {
	return len(*h)
}

func (h *headHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *headHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *headHeap) Push(x interface{}) {

	*h = append(*h, x.(*ListNode))
}

func (h *headHeap) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := ListNode{}
	end := &res
	hp := headHeap{}
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(&hp, lists[i])
	}

	for hp.Len() > 0 {
		e := heap.Pop(&hp).(*ListNode)
		end.Next = e
		end = e
		if e.Next != nil {
			heap.Push(&hp, e.Next)
		}
	}

	return res.Next
}

func createList(nodes []ListNode, shift int) *ListNode {
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Val = i + shift
		nodes[i].Next = &nodes[i+1]
	}
	nodes[len(nodes)-1].Val = len(nodes) - 1 + shift
	return &nodes[0]
}

func main() {
	l1 := createList(make([]ListNode, 2), 3)
	l2 := createList(make([]ListNode, 4), 2)
	l3 := createList(make([]ListNode, 3), 1)
	r := mergeKLists([]*ListNode{l1, l2, l3})
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
	mergeKLists([]*ListNode{nil})
}
