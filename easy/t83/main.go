package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return head
}
