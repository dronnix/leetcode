package main

import (
	"fmt"
	"math"
)

// Task: https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt, math.MaxInt)
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}

func main() {
	fmt.Println(isValidBST(nil))
}
