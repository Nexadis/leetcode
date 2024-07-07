package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	return max(l, r)
}
