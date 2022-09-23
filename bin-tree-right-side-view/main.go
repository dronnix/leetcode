package main

// The task: https://leetcode.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0, 100)
	walk(root, &res, 0)
	return res
}

func walk(root *TreeNode, res *[]int, lvl int) {
	if root == nil {
		return
	}
	if len(*res) == lvl {
		*res = append(*res, root.Val)
	}

	walk(root.Right, res, lvl+1)
	walk(root.Left, res, lvl+1)
}

func main() {
	rightSideView(nil)
}
