package main

import "fmt"

// The task: https://leetcode.com/problems/sort-list/solution/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	result := ListNode{Val: -100001}
	for head != nil {
		ins := &result
		for ins.Next != nil {
			if head.Val < ins.Next.Val {
				break
			}
			ins = ins.Next
		}
		h := head
		head = head.Next
		ins.Next, h.Next = h, ins.Next
	}
	return result.Next
}

func main() {
	l := make([]ListNode, 4)
	l[0] = ListNode{4, &l[1]}
	l[1] = ListNode{1, &l[2]}
	l[2] = ListNode{2, &l[3]}
	l[3] = ListNode{3, nil}
	r := sortList(&l[0])
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
