package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var reversed *ListNode
	for slow != nil {
		slow.Next, reversed, slow = reversed, slow, slow.Next
	}

	for reversed.Next != nil {
		head.Next, head = reversed, head.Next
		reversed.Next, reversed = head, reversed.Next
	}
}

func main() {
	nodes := make([]ListNode, 5)
	for i := 0; i < 4; i++ {
		nodes[i].Val = i + 1
		nodes[i].Next = &nodes[i+1]
	}
	nodes[4].Val = 5

	reorderList(&nodes[0])

	for p := &nodes[0]; p != nil; {
		fmt.Println(p.Val)
		p = p.Next
	}

}
