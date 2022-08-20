package main

import (
	"fmt"
	"testing"
)

func getTree() *TreeNode {
	nodes := make([]TreeNode, 5)
	for i := range nodes {
		nodes[i].Val = i + 1
	}
	nodes[2].Right = &nodes[1]
	nodes[2].Left = &nodes[3]
	nodes[1].Right = &nodes[0]
	nodes[3].Left = &nodes[4]
	return &nodes[2]
}

func TestSerialization(t *testing.T) {
	c := Constructor()
	s := c.serialize(getTree()) // nolint:ifshort
	if s != "3,2,1,65535,65535,65535,4,65535,5,65535,65535" {
		t.Fail()
	}
}

func TestDeserialization(t *testing.T) {
	s1 := "3,2,1,65535,65535,65535,4,65535,5,65535,65535"
	c := Constructor()
	r := c.deserialize(s1)
	s2 := c.serialize(r) // nolint:ifshort
	if s1 != s2 {
		fmt.Println(s1, s2)
		t.Fail()
	}
}
