package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoList(list1.Next, list2)
	} else {
		list2.Next = mergeTwoList(list1, list2.Next)
	}
	return &ListNode{}
}

// Function to print the linked list
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println() // for a new line after printing the list
}

func main() {
	list11 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		},
	}

	list21 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	}

	printList(&list11)
	printList(&list21)
}
