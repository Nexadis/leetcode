package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	cnt := make(map[*ListNode]int)
	for head != nil {
		cnt[head]++
		if cnt[head] > 1 {
			return true
		}
		head = head.Next
	}
	return false
}
