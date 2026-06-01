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
			Val: 3,
			Next: &MergeNode{
				Val: 5,
				Next: &MergeNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	HeadTwo = &MergeNode{
		Val: 2,
		Next: &MergeNode{
			Val: 4,
			Next: &MergeNode{
				Val: 6,
				Next: &MergeNode{
					Val:  8,
					Next: nil,
				},
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
	merged := &MergeNode{}

	if currentOne.Val <= currentTwo.Val {
		merged = currentOne
		currentOne = currentOne.Next
	} else {
		merged = currentTwo
		currentTwo = currentTwo.Next
	}
	head := merged

	count := 0

	for merged != nil {
		if currentOne.Val <= currentTwo.Val {
			merged.Next = currentOne
			currentOne = currentOne.Next
		} else {
			merged.Next = currentTwo
			currentTwo = currentTwo.Next
		}
		merged = merged.Next
		count++
	}
	return head
}
