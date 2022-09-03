package main

import "fmt"

// The task: https://leetcode.com/problems/sort-list/solution/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	lists := divide(head)
	return mergeAll(lists)
}

func divide(head *ListNode) []*ListNode {
	result := make([]*ListNode, 0)
	for head != nil {
		result = append(result, head)
		cur := head
		head = head.Next
		cur.Next = nil
	}
	if len(result)%2 > 0 {
		result = append(result, nil)
	}
	return result
}

func mergeAll(lists []*ListNode) *ListNode {
	for {
		nl := make([]*ListNode, 0, len(lists)/2+1)
		for i := 0; i < len(lists)-1; i += 2 {
			nl = append(nl, merge(lists[i], lists[i+1]))
		}
		if len(nl) == 1 {
			return nl[0]
		}
		if len(nl)%2 > 0 {
			nl = append(nl, nil)
		}
		lists = nl
	}
}

func merge(l1, l2 *ListNode) *ListNode {
	fake := ListNode{}
	ins := &fake
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ins.Next = l1
			l1 = l1.Next
		} else {
			ins.Next = l2
			l2 = l2.Next
		}
		ins = ins.Next
	}
	if l1 != nil {
		ins.Next = l1
	}
	if l2 != nil {
		ins.Next = l2
	}
	return fake.Next
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
