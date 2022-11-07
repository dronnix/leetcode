package diam_bin_tree

// The task: https://leetcode.com/problems/diameter-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int { //nolint:unused
	m := 0
	measureDiameter(root, &m)
	return m
}

func measureDiameter(root *TreeNode, max *int) (path int) { //nolint:unused
	if root == nil {
		return 0
	}

	rp := 0
	if root.Right != nil {
		rp = measureDiameter(root.Right, max) + 1
	}

	lp := 0
	if root.Left != nil {
		lp = measureDiameter(root.Left, max) + 1
	}

	if lp+rp > *max {
		*max = lp + rp
	}

	if lp > rp {
		return lp
	}
	return rp
}
