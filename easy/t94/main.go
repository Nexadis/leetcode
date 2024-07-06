package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func inorderTraversal(root *TreeNode) []int {
	return walk(root)
}

func walk(t *TreeNode) []int {
	if t == nil {
		return nil
	}
	l := walk(t.Left)
	m := t.Val
	r := walk(t.Right)
	res := make([]int, 0, len(l)+len(r)+1)
	res = append(res, l...)
	res = append(res, m)
	res = append(res, r...)
	return res
}
