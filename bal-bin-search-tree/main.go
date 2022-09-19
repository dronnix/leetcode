package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0, 256)
	walk(root, &arr)
	return split(arr)

}

func walk(root *TreeNode, arr *[]int) {
	if root.Left != nil {
		walk(root.Left, arr)
	}

	*arr = append(*arr, root.Val)

	if root.Right != nil {
		walk(root.Right, arr)
	}
}

func split(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	m := len(arr) / 2
	return &TreeNode{
		Val:   arr[m],
		Left:  split(arr[:m]),
		Right: split(arr[m+1:]),
	}
}

func main() {
	tree := make([]TreeNode, 7)
	tree[0].Val = 1
	tree[0].Right = &tree[1]
	tree[1].Val = 2
	tree[1].Right = &tree[2]
	tree[2].Val = 3
	tree[2].Right = &tree[3]
	tree[3].Val = 4
	tree[3].Right = &tree[4]
	tree[4].Val = 5
	tree[4].Right = &tree[5]
	tree[5].Val = 6
	tree[5].Right = &tree[6]
	tree[6].Val = 7
	r := balanceBST(&tree[0])
	fmt.Println(r.Val)
}
