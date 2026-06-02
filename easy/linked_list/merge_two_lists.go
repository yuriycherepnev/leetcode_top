package main

import "fmt"

type MergeNode struct {
	Val  int
	Next *MergeNode
}

var (
	HeadOne = &MergeNode{
		Val: 1,
		Next: &MergeNode{
			Val: 2,
			Next: &MergeNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	HeadTwo = &MergeNode{
		Val: 1,
		Next: &MergeNode{
			Val: 3,
			Next: &MergeNode{
				Val:  4,
				Next: nil,
			},
		},
	}
)

func main() {
	merged := mergeTwoLists(HeadOne, HeadTwo)

	for merged != nil {
		fmt.Println(merged)
		merged = merged.Next
	}
}

func mergeTwoLists(list1 *MergeNode, list2 *MergeNode) *MergeNode {
	currentOne := list1
	currentTwo := list2
	prev := &MergeNode{}
	head := prev

	for currentOne != nil && currentTwo != nil {
		if currentOne.Val <= currentTwo.Val {
			prev.Next = currentOne
			currentOne = currentOne.Next
		} else {
			prev.Next = currentTwo
			currentTwo = currentTwo.Next
		}
		prev = prev.Next
	}

	if currentOne != nil {
		prev.Next = currentOne
	}
	if currentTwo != nil {
		prev.Next = currentTwo
	}

	return head.Next
}
