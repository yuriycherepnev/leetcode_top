package main

import "fmt"

type EndNode struct {
	Val  int
	Next *EndNode
}

type EndList struct {
	head *EndNode
}

func (l *EndList) PushForward(val int) {
	newNode := &EndNode{
		Val: val,
	}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.Next = l.head
	l.head = newNode
}

func main() {
	list := EndList{}
	list.PushForward(5)
	list.PushForward(4)
	list.PushForward(3)
	list.PushForward(2)
	list.PushForward(1)
	removeNthFromEnd(list.head, 2)
	current := list.head
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

func removeNthFromEnd(head *EndNode, n int) *EndNode {
	node := &EndNode{
		Next: head,
	}
	ahead := node
	base := node

	for i := 0; i < n; i++ {
		ahead = ahead.Next
	}

	for ahead.Next != nil {
		ahead = ahead.Next
		base = base.Next
	}
	base.Next = base.Next.Next

	return node.Next
}
