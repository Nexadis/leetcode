package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := preorderTraversal(root.Left)
	nums = append([]int{root.Val}, nums...)
	return append(nums, preorderTraversal(root.Right)...)
}
