package main

import "fmt"

// The task: https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func main() {
	nodes := make([]TreeNode, 3)
	for i := range nodes {
		nodes[i].Val = i
	}
	nodes[1].Left = &nodes[0]
	nodes[1].Right = &nodes[2]
	r := invertTree(&nodes[1])
	fmt.Println(r.Left.Val, r.Right.Val)
}
