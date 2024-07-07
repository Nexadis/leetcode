package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return checkSymmetric(l.Left, r.Right) && checkSymmetric(r.Left, l.Right)
}
