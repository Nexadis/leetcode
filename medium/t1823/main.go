package main

func findTheWinner(n int, k int) int {
	friends := make([]int, n)
	for i := range friends {
		friends[i] = i + 1
	}
	head := NewList(friends)
	l := n
	for l > 1 {
		head = fwdDel(head, k-1)
		l--
	}
	return head.Val
}

func NewList(nums []int) *Node {
	head := &Node{
		Val: nums[0],
	}
	tail := head
	for i := 1; i < len(nums); i++ {
		n := &Node{
			Prev: tail,
			Val:  nums[i],
			Next: head,
		}
		tail.Next = n
		tail = n
	}
	head.Prev = tail
	return head
}

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

func fwdDel(n *Node, steps int) (next *Node) {
	for i := 0; i < steps; i++ {
		n = n.Next
	}
	return del(n)
}

func del(n *Node) (next *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	return n.Next
}
