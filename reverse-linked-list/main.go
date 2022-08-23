package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = tail
		tail = cur
	}
	return tail
}

func reverseListRecursive(head *ListNode) *ListNode { // nolint:unused,deadcode
	return reverse(head, nil)
}

func reverse(head, tail *ListNode) (newHead *ListNode) { // nolint:unused,deadcode
	if head == nil {
		return tail
	}
	tmp := head.Next
	head.Next = tail
	return reverse(tmp, head)
}

func main() {
	nodes := make([]ListNode, 5)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Val = i
		nodes[i].Next = &nodes[i+1]
	}
	l := reverseList(&nodes[0])
	for ; l != nil; l = l.Next {
		fmt.Println(l.Val)
	}
}
