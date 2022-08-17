package main

// Task: https://leetcode.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	walker, runner := head, head
	for runner != nil {
		if runner.Next == nil {
			return walker
		}
		runner = runner.Next.Next
		walker = walker.Next
	}
	return walker
}

func main() {

}
