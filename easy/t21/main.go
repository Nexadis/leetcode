package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Next *ListNode
	Val  int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var begin *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list2.Val < list1.Val {
		begin = list2
		list2 = list2.Next
	} else {
		begin = list1
		list1 = list1.Next
	}
	cur := begin

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		cur.Next = list1
		cur = cur.Next
		list1 = list1.Next
	}
	for list2 != nil {
		cur.Next = list2
		cur = cur.Next
		list2 = list2.Next
	}
	return begin
}
