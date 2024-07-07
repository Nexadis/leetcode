package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	root := &TreeNode{
		Val: nums[m],
	}
	l := sortedArrayToBST(nums[:m])
	r := (*TreeNode)(nil)
	if m+1 < len(nums) {
		r = sortedArrayToBST(nums[m+1:])
	}
	root.Left = l
	root.Right = r
	return root
}
