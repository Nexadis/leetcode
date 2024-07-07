package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Right == nil && root.Left == nil {
		return true
	}
	targetSum -= root.Val
	has := hasPathSum(root.Left, targetSum)
	has = has || hasPathSum(root.Right, targetSum)
	return has
}
