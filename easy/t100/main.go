package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	res := true
	if p == nil && q == nil {
		return res
	}
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	res = res && isSameTree(p.Left, q.Left)
	res = res && isSameTree(p.Right, q.Right)

	return res
}
