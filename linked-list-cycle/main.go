package main

import "fmt"

// Task: https://leetcode.com/problems/linked-list-cycle/submissions/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	walker, runner := head.Next, head.Next
	for {
		walker = walker.Next
		if runner == nil || runner.Next == nil {
			return false
		}
		runner = runner.Next.Next
		if runner == walker {
			return true
		}
	}
}

func main() {
	nodes := make([]ListNode, 5)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Val = i
		nodes[i].Next = &nodes[i+1]
	}
	fmt.Println(hasCycle(&nodes[0]))
	nodes[4].Next = &nodes[0]
	fmt.Println(hasCycle(&nodes[0]))
}
