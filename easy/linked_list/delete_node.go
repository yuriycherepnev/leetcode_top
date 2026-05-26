package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func (l *List) PushForward(val int) *ListNode {
	newNode := &ListNode{
		Val: val,
	}
	if l.Head == nil {
		l.Head = newNode
		return newNode
	}
	newNode.Next = l.Head
	l.Head = newNode

	return newNode
}

func main() {
	list := List{}

	list.PushForward(10)
	list.PushForward(20)
	list.PushForward(30)
	list.PushForward(40)

	current := list.Head

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next

	fmt.Println(node)
}

func deleteNodeTwo(node *ListNode) {
	*node = *node.Next
}

/*
node - это указатель на структуру ListNode
*node — это сама структура (её содержимое: поля Val и Next)

node.Next — это указатель на следующий узел.
*node.Next — это структура следующего узла (его поля Val и Next)

Что делает присваивание *node = *node.Next?
Оно копирует все поля следующего узла в текущий узел.
То есть текущий узел становится точной копией следующего узла.
*/
