package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	fake, sum := ListNode{}, 0
	for last := &fake; l1 != nil || l2 != nil || sum != 0; last = last.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		last.Next = &ListNode{sum % 10, nil}
		sum /= 10
	}
	return fake.Next
}

func main() {
	n1, n2 := make([]ListNode, 2), make([]ListNode, 2)
	n1[0].Val, n1[1].Val = 7, 6
	n2[0].Val, n2[1].Val = 3, 3
	n1[0].Next = &n1[1]
	n2[0].Next = &n2[1]

	r := addTwoNumbers(&n1[0], &n2[0])
	for ; r != nil; r = r.Next {
		fmt.Println(r.Val)
	}
}
