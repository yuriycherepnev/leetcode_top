package main

import "fmt"

type ReverseNode struct {
	Val  int
	Next *ReverseNode
}

func main() {
	nodeSix := &ReverseNode{Val: 6, Next: nil}
	nodeFive := &ReverseNode{Val: 5, Next: nodeSix}
	nodeFour := &ReverseNode{Val: 4, Next: nodeFive}
	nodeThree := &ReverseNode{Val: 3, Next: nodeFour}
	nodeTwo := &ReverseNode{Val: 2, Next: nodeThree}
	nodeOne := &ReverseNode{Val: 1, Next: nodeTwo}

	head := nodeOne

	current := reverseList(head)

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}

func reverseList(head *ReverseNode) *ReverseNode {
	var prev *ReverseNode // nil
	current := head       // ReverseNode{Val: 1, Next: nodeTwo}

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

/*
var prev *ReverseNode — это объявление переменной-указателя,
которая изначально указывает на nil

В Go любое объявление переменной ссылочного типа (указатель, slice, map, channel, функция, интерфейс)
без явной инициализации получает нулевое значение (zero value), которое равно nil.


*/
