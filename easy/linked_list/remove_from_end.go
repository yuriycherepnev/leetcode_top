package main

import "fmt"

type EndNode struct {
	Value int
	Next  *EndNode
}

type EndList struct {
	head *EndNode
}

func (l *EndList) PushForward(val int) {
	newNode := &EndNode{
		Value: val,
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

	list.PushForward(10)
	list.PushForward(20)
	list.PushForward(30)
	list.PushForward(40)

	result := removeNthFromEnd(&list, 2)

	current := result.head
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}

}

func removeNthFromEnd(head *EndList, n int) *EndList {
	current := head.head
	listCount := 1

	for current != nil {
		if listCount == n {
			current.Value = current.Next.Value
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
	return head
}
