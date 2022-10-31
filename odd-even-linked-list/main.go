package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	fTmp, sTmp := ListNode{}, ListNode{}
	fEnd, sEnd := &fTmp, &sTmp

	for head != nil {
		fEnd.Next = head
		fEnd = head

		head = head.Next
		if head == nil {
			break
		}

		sEnd.Next = head
		sEnd = head

		head = head.Next
	}

	fEnd.Next = sTmp.Next
	sEnd.Next = nil
	return fTmp.Next
}

func main() {
	elems := make([]ListNode, 5)
	for i := 0; i < 4; i++ {
		elems[i].Val = i
		elems[i].Next = &elems[i+1]
	}
	elems[4] = ListNode{4, nil}

	r := oddEvenList(&elems[0])
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
