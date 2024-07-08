package main

func findTheWinner(n int, k int) int {
	head := NewList(n)
	l := n
	for l > 1 {
		head = fwdDel(head, k-1)
		l--
	}
	return head.Val
}

func NewList(size int) *Node {
	head := &Node{
		Val: 1,
	}
	tail := head
	for i := 1; i < size; i++ {
		n := &Node{
			Prev: tail,
			Val:  i + 1,
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
