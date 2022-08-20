package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Task: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	buff := &strings.Builder{}
	buff.WriteString(fmt.Sprintf("%d", root.Val))
	buff = serialize(root.Right, buff)
	buff = serialize(root.Left, buff)
	return buff.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	parts := strings.Split(data, ",")
	nums := make([]int, 0, len(parts))
	for i := range parts {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	root := &TreeNode{Val: nums[0]}
	deserialize(root, nums, 1)

	return root
}

const null = 65535

func deserialize(parent *TreeNode, nums []int, pos int) int {
	if nums[pos] == null {
		parent.Right = nil
		pos++
	} else {
		parent.Right = &TreeNode{Val: nums[pos]}
		pos = deserialize(parent.Right, nums, pos+1)
	}
	if nums[pos] == null {
		parent.Left = nil
		pos++
	} else {
		parent.Left = &TreeNode{Val: nums[pos]}
		pos = deserialize(parent.Left, nums, pos+1)
	}
	return pos
}

func serialize(node *TreeNode, b *strings.Builder) *strings.Builder {
	if node == nil {
		b.WriteString(",65535")
		return b
	}
	b.WriteString(",")
	b.WriteString(strconv.Itoa(node.Val))
	b = serialize(node.Right, b)
	b = serialize(node.Left, b)
	return b
}

func main() {

}
