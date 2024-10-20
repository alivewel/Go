package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	for list1 != nil || list2 != nil {
		fmt.Println("!")
		if list1 != nil {
			res.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else if list2 != nil {
			res.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
	}
	return res.Next
}

func main() {
	// Список: 1->2->4
	first := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// Список: 1->3->4
	second := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	fmt.Println(mergeTwoLists(first, second)) // Список: 1->1->2->3->4->4
}
