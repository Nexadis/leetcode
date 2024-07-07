package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if r*l != 0 {
		return min(l, r) + 1
	}
	return l + r + 1
}
