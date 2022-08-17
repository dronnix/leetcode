package main

import (
	"testing"
)

func createList(n int) []ListNode {
	nodes := make([]ListNode, n)
	for i := 1; i < n; i++ {
		nodes[i].Val = i
		nodes[i-1].Next = &nodes[i]
	}
	return nodes
}

func Test_middleNode_Odd(t *testing.T) {
	t.Parallel()
	nodes := createList(5)
	if mid := middleNode(&nodes[0]); mid != &nodes[2] {
		t.Fail()
	}
}

func Test_middleNode_Even(t *testing.T) {
	t.Parallel()
	nodes := createList(4)
	if mid := middleNode(&nodes[0]); mid != &nodes[2] {
		t.Fail()
	}
}

func Test_middleNode_One(t *testing.T) {
	t.Parallel()
	node := ListNode{Val: 3}
	if mid := middleNode(&node); mid != &node {
		t.Fail()
	}
}

func Test_middleNode_Zero(t *testing.T) {
	t.Parallel()
	if mid := middleNode(nil); mid != nil {
		t.Fail()
	}
}
