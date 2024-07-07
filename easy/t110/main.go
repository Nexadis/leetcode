package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	l := depth(root.Left)
	r := depth(root.Right)
	return max(l, r)-min(l, r) <= 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	return max(l, r) + 1
}
